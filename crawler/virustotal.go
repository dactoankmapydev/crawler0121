package crawler

import (
	"encoding/json"
	"fmt"
	"ioc-provider/helper"
	"ioc-provider/model"
	"ioc-provider/repository"
	"strings"
)

type VirustotalResult struct {
	Data []struct {
		Attributes struct {
			Names               []string `json:"names"`
			Md5                 string   `json:"md5"`
			Sha1                string   `json:"sha1"`
			Sha256              string   `json:"sha256"`
			Tags                []string `json:"tags"`
			FirstSubmissionDate int      `json:"first_submission_date"`
			Exiftool            struct {
				FileType string `json:"FileType"`
			} `json:"exiftool"`
			LastAnalysisResults map[string]map[string]string `json:"last_analysis_results"`
		} `json:"attributes"`
		ContextAttributes struct {
			NotificationDate int `json:"notification_date"`
		} `json:"context_attributes"`
		ID string `json:"id"`
	} `json:"data"`
	Meta struct {
		Cursor string `json:"cursor"`
	} `json:"meta"`
}

func LiveHunting(repo repository.IocRepo) {
	existsPost := repo.ExistsIndex(model.IndexNameSample)
	if !existsPost {
		repo.CreateIndex(model.MappingSample, model.MappingSample)
	}
	sampleList := make([]model.Sample, 0)
	cursor := []string{""}
	for len(cursor) > 0 {
		pathAPI := fmt.Sprintf("https://www.virustotal.com/api/v3/intelligence/hunting_notification_files?cursor=%s", cursor[0]+"&limit=40")
		fmt.Println("pathAPI->", pathAPI)
		body, err := helper.HttpClient.GetRequestVirustotal(pathAPI)
		if err != nil {
			return
		}
		var virustotalResult VirustotalResult
		json.Unmarshal(body, &virustotalResult)
		if virustotalResult.Meta.Cursor != "" {
			cursor[0] = virustotalResult.Meta.Cursor
			for i, item := range virustotalResult.Data {
				pointAv := virustotalResult.enginesPoint(i)
				if pointAv >= 13 {
					sample := model.Sample{
						Name:             strings.Join(item.Attributes.Names, ", "),
						Sha256:           item.Attributes.Sha256,
						Sha1:             item.Attributes.Sha1,
						Md5:              item.Attributes.Md5,
						Tags:             item.Attributes.Tags,
						FirstSubmit:      item.Attributes.FirstSubmissionDate,
						NotificationDate: item.ContextAttributes.NotificationDate,
						FileType:         item.Attributes.Exiftool.FileType,
						EnginesDetected:  virustotalResult.enginesDetected(i),
						Detected:         len(virustotalResult.enginesDetected(i)),
						Point:            virustotalResult.enginesPoint(i),
					}
					sampleList = append(sampleList, sample)
					existsID := repo.ExistsDoc(model.IndexNameSample, sample.Sha256)
					if existsID {
						break
					} else {
						success := repo.InsertIndex(model.IndexNameSample, sample.Sha256, sample)
						if !success {
							return
						}
					}
					fmt.Println("sample->", sample)
				}
			}
		} else {
			cursor = cursor[:0]
		}
	}
	fmt.Println("len listSample->", len(sampleList))
}

// Lọc ra loại engines detected
func enginesTypeDetected(enginesType []string, enginesTypeClear []string) []string {
	var typeDetected []string
	for i := 0; i < len(enginesType); i++ {
		var isExit bool
		for j := 0; j < len(enginesTypeClear); j++ {
			if enginesType[i] == enginesTypeClear[j] {
				isExit = true
				break
			}
		}
		if isExit != true {
			typeDetected = append(typeDetected, enginesType[i])
		}
	}
	return typeDetected
}

// Hợp nhất tên engines và kiểu engines detected thành một map
func merge(avName []string, avType []string) map[string]string {
	avMap := make(map[string]string)
	for i := 0; i < len(avName); i++ {
		for j := 0; j < len(avType); j++ {
			avMap[avName[i]] = avType[i]
		}
	}
	return avMap
}

// Lọc ra tên engines detected
func nameEnginesDetected(typeDetected []string, engines map[string]string) []string {
	var nameDetected []string
	for i := 0; i < len(typeDetected); i++ {
		for nameEngines, typeEngines := range engines {
			if typeDetected[i] == typeEngines {
				nameDetected = append(nameDetected, nameEngines)
			}
		}
		break
	}
	return nameDetected
}

// Tính tổng điểm cho engines có tên nằm trong enginesHash
func point(enginesDetected []string) int {
	enginesHash := map[string]int{
		"Ad-Aware":                 1,
		"AegisLab":                 1,
		"ALYac":                    2,
		"Antiy-AVL":                1,
		"Arcabit":                  1,
		"Avast":                    3,
		"AVG":                      2,
		"Avira":                    1,
		"Baidu":                    2,
		"BitDefender":              3,
		"CAT-QuickHeal":            1,
		"Comodo":                   2,
		"Cynet":                    1,
		"Cyren":                    1,
		"DrWeb":                    1,
		"Emsisoft":                 2,
		"eScan":                    2,
		"ESET-NOD32":               3,
		"F-Secure":                 2,
		"FireEye":                  3,
		"Fortinet":                 3,
		"GData":                    1,
		"Ikarus":                   2,
		"Kaspersky":                3,
		"MAX":                      1,
		"McAfee":                   3,
		"Microsoft":                3,
		"Panda":                    2,
		"Qihoo-360":                2,
		"Rising":                   1,
		"Sophos":                   2,
		"TrendMicro":               3,
		"TrendMicro-HouseCall":     1,
		"ZoneAlarm by Check Point": 1,
		"Zoner":                    1,
		"AhnLab - V3":              1,
		"BitDefenderTheta":         2,
		"Bkav":                     1,
		"ClamAV":                   3,
		"CMC":                      1,
		"Gridinsoft":               1,
		"Jiangmin":                 1,
		"K7AntiVirus":              1,
		"K7GW":                     1,
		"Kingsoft":                 1,
		"Malwarebytes":             3,
		"MaxSecure":                1,
		"McAfee - GW - Edition":    3,
		"NANO - Antivirus":         1,
		"Sangfor Engine Zero":      1,
		"SUPERAntiSpyware":         1,
		"Symantec":                 3,
		"TACHYON":                  1,
		"Tencent":                  2,
		"TotalDefense":             1,
		"VBA32":                    2,
		"VIPRE":                    1,
		"ViRobot":                  1,
		"Yandex":                   3,
		"Zillya":                   1,
		"Acronis":                  3,
		"Alibaba":                  2,
		"SecureAge APEX":           1,
		"Avast - Mobile":           2,
		"BitDefenderFalx":          3,
		"CrowdStrike Falcon":       3,
		"Cybereason":               3,
		"Cylance":                  2,
		"eGambit":                  1,
		"Elastic":                  1,
		"Palo Alto Networks":       2,
		"SentinelOne (Static ML)":  1,
		"Symantec Mobile Insight":  3,
		"Trapmine":                 1,
		"Trustlook":                1,
		"Webroot":                  1,
	}
	var total int = 0
	for i := 0; i < len(enginesDetected); i++ {
		for nameEngines, pointEngines := range enginesHash {
			if nameEngines == enginesDetected[i] {
				total += pointEngines
			}
		}
	}
	return total
}

// Danh sách engines detected
func (vr VirustotalResult) enginesDetected(i int) []string {
	enginesType := make([]string, 0)
	enginesName := make([]string, 0)
	enginesTypeClear := []string{"confirmed-timeout", "undetected", "timeout", "type-unsupported", "failure"}
	for index, item := range vr.Data {
		if index == i {
			totalEngines := item.Attributes.LastAnalysisResults
			for avName, avType := range totalEngines {
				enginesName = append(enginesName, avName)
				enginesType = append(enginesType, avType["category"])
			}
		}
	}
	detect := enginesTypeDetected(enginesType, enginesTypeClear)
	engines := merge(enginesName, enginesType)
	return nameEnginesDetected(detect, engines)
}

// Tính điểm engines
func (vr VirustotalResult) enginesPoint(i int) int {
	return point(vr.enginesDetected(i))
}
