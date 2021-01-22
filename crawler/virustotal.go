package crawler

import (
	"encoding/json"
	"fmt"
	"ioc-provider/helper"
	"ioc-provider/model"
	"ioc-provider/repository"
	"strings"
)

//type Go struct {
//	Data []GoData `json:"data"`
//	Meta GoMeta `json:"meta"`
//}
//
//type GoExtensions struct {
//	Bin int `json:"bin"`
//	XML int `json:"xml"`
//}
//
//type GoFileTypes struct {
//	JPG int `json:"JPG"`
//	MicrosoftOffice int `json:"Microsoft Office"`
//	XML int `json:"XML"`
//}
//
//type GoBundleInfo struct {
//	Extensions GoExtensions `json:"extensions"`
//	FileTypes GoFileTypes `json:"file_types"`
//	HighestDatetime string `json:"highest_datetime"`
//	LowestDatetime string `json:"lowest_datetime"`
//	NumChildren int `json:"num_children"`
//	Type string `json:"type"`
//	UncompressedSize int `json:"uncompressed_size"`
//}
//
//type GoExiftool struct {
//	AppVersion string `json:"AppVersion"`
//	Application string `json:"Application"`
//	Characters string `json:"Characters"`
//	CharactersWithSpaces string `json:"CharactersWithSpaces"`
//	CreateDate string `json:"CreateDate"`
//	Creator string `json:"Creator"`
//	DocSecurity string `json:"DocSecurity"`
//	FileType string `json:"FileType"`
//	FileTypeExtension string `json:"FileTypeExtension"`
//	HyperlinksChanged string `json:"HyperlinksChanged"`
//	LastModifiedBy string `json:"LastModifiedBy"`
//	Lines string `json:"Lines"`
//	LinksUpToDate string `json:"LinksUpToDate"`
//	MIMEType string `json:"MIMEType"`
//	ModifyDate string `json:"ModifyDate"`
//	Pages string `json:"Pages"`
//	Paragraphs string `json:"Paragraphs"`
//	RevisionNumber string `json:"RevisionNumber"`
//	ScaleCrop string `json:"ScaleCrop"`
//	SharedDoc string `json:"SharedDoc"`
//	Template string `json:"Template"`
//	TotalEditTime string `json:"TotalEditTime"`
//	Words string `json:"Words"`
//	ZipBitFlag string `json:"ZipBitFlag"`
//	ZipCRC string `json:"ZipCRC"`
//	ZipCompressedSize string `json:"ZipCompressedSize"`
//	ZipCompression string `json:"ZipCompression"`
//	ZipFileName string `json:"ZipFileName"`
//	ZipModifyDate string `json:"ZipModifyDate"`
//	ZipRequiredVersion string `json:"ZipRequiredVersion"`
//	ZipUncompressedSize string `json:"ZipUncompressedSize"`
//}
//
//type GoALYac struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoAPEX struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoAcronis struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoAdAware struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoAegisLab struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoAhnLabV3 struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoAlibaba struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoAntiyAVL struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoArcabit struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoAvast struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoAvastMobile struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoAvira struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoBaidu struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoBitDefender struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoBitDefenderFalx struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoBitDefenderTheta struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoBkav struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoCATQuickHeal struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoCMC struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoClamAV struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoComodo struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoCrowdStrike struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoCybereason struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoCylance struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoCynet struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoCyren struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoDrWeb struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoESETNOD32 struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoElastic struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoEmsisoft struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoFSecure struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoFireEye struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoFortinet struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoGData struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoGridinsoft struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoIkarus struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoJiangmin struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoK7AntiVirus struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoK7GW struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoKaspersky struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoKingsoft struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoMAX struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoMalwarebytes struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoMaxSecure struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoMcAfee struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoMcAfeeGWEdition struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoMicroWorldEScan struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoMicrosoft struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoNANOAntivirus struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoPaloalto struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoPanda struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoQihoo360 struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoRising struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoSUPERAntiSpyware struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoSangfor struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoSentinelOne struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoSophos struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoSymantec struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoSymantecMobileInsight struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoTACHYON struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoTencent struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoTotalDefense struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoTrapmine struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoTrendMicro struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoTrendMicroHouseCall struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoTrustlook struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoVBA32 struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoVIPRE struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoViRobot struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoWebroot struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoYandex struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result string `json:"result"`
//}
//
//type GoZillya struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoZoneAlarm struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoZoner struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion string `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoEGambit struct {
//	Category string `json:"category"`
//	EngineName string `json:"engine_name"`
//	EngineUpdate string `json:"engine_update"`
//	EngineVersion interface{} `json:"engine_version"`
//	Method string `json:"method"`
//	Result interface{} `json:"result"`
//}
//
//type GoLastAnalysisResults struct {
//	ALYac GoALYac `json:"ALYac"`
//	APEX GoAPEX `json:"APEX"`
//	Acronis GoAcronis `json:"Acronis"`
//	AdAware GoAdAware `json:"Ad-Aware"`
//	AegisLab GoAegisLab `json:"AegisLab"`
//	AhnLabV3 GoAhnLabV3 `json:"AhnLab-V3"`
//	Alibaba GoAlibaba `json:"Alibaba"`
//	AntiyAVL GoAntiyAVL `json:"Antiy-AVL"`
//	Arcabit GoArcabit `json:"Arcabit"`
//	Avast GoAvast `json:"Avast"`
//	AvastMobile GoAvastMobile `json:"Avast-Mobile"`
//	Avira GoAvira `json:"Avira"`
//	Baidu GoBaidu `json:"Baidu"`
//	BitDefender GoBitDefender `json:"BitDefender"`
//	BitDefenderFalx GoBitDefenderFalx `json:"BitDefenderFalx"`
//	BitDefenderTheta GoBitDefenderTheta `json:"BitDefenderTheta"`
//	Bkav GoBkav `json:"Bkav"`
//	CATQuickHeal GoCATQuickHeal `json:"CAT-QuickHeal"`
//	CMC GoCMC `json:"CMC"`
//	ClamAV GoClamAV `json:"ClamAV"`
//	Comodo GoComodo `json:"Comodo"`
//	CrowdStrike GoCrowdStrike `json:"CrowdStrike"`
//	Cybereason GoCybereason `json:"Cybereason"`
//	Cylance GoCylance `json:"Cylance"`
//	Cynet GoCynet `json:"Cynet"`
//	Cyren GoCyren `json:"Cyren"`
//	DrWeb GoDrWeb `json:"DrWeb"`
//	ESETNOD32 GoESETNOD32 `json:"ESET-NOD32"`
//	Elastic GoElastic `json:"Elastic"`
//	Emsisoft GoEmsisoft `json:"Emsisoft"`
//	FSecure GoFSecure `json:"F-Secure"`
//	FireEye GoFireEye `json:"FireEye"`
//	Fortinet GoFortinet `json:"Fortinet"`
//	GData GoGData `json:"GData"`
//	Gridinsoft GoGridinsoft `json:"Gridinsoft"`
//	Ikarus GoIkarus `json:"Ikarus"`
//	Jiangmin GoJiangmin `json:"Jiangmin"`
//	K7AntiVirus GoK7AntiVirus `json:"K7AntiVirus"`
//	K7GW GoK7GW `json:"K7GW"`
//	Kaspersky GoKaspersky `json:"Kaspersky"`
//	Kingsoft GoKingsoft `json:"Kingsoft"`
//	MAX GoMAX `json:"MAX"`
//	Malwarebytes GoMalwarebytes `json:"Malwarebytes"`
//	MaxSecure GoMaxSecure `json:"MaxSecure"`
//	McAfee GoMcAfee `json:"McAfee"`
//	McAfeeGWEdition GoMcAfeeGWEdition `json:"McAfee-GW-Edition"`
//	MicroWorldEScan GoMicroWorldEScan `json:"MicroWorld-eScan"`
//	Microsoft GoMicrosoft `json:"Microsoft"`
//	NANOAntivirus GoNANOAntivirus `json:"NANO-Antivirus"`
//	Paloalto GoPaloalto `json:"Paloalto"`
//	Panda GoPanda `json:"Panda"`
//	Qihoo360 GoQihoo360 `json:"Qihoo-360"`
//	Rising GoRising `json:"Rising"`
//	SUPERAntiSpyware GoSUPERAntiSpyware `json:"SUPERAntiSpyware"`
//	Sangfor GoSangfor `json:"Sangfor"`
//	SentinelOne GoSentinelOne `json:"SentinelOne"`
//	Sophos GoSophos `json:"Sophos"`
//	Symantec GoSymantec `json:"Symantec"`
//	SymantecMobileInsight GoSymantecMobileInsight `json:"SymantecMobileInsight"`
//	TACHYON GoTACHYON `json:"TACHYON"`
//	Tencent GoTencent `json:"Tencent"`
//	TotalDefense GoTotalDefense `json:"TotalDefense"`
//	Trapmine GoTrapmine `json:"Trapmine"`
//	TrendMicro GoTrendMicro `json:"TrendMicro"`
//	TrendMicroHouseCall GoTrendMicroHouseCall `json:"TrendMicro-HouseCall"`
//	Trustlook GoTrustlook `json:"Trustlook"`
//	VBA32 GoVBA32 `json:"VBA32"`
//	VIPRE GoVIPRE `json:"VIPRE"`
//	ViRobot GoViRobot `json:"ViRobot"`
//	Webroot GoWebroot `json:"Webroot"`
//	Yandex GoYandex `json:"Yandex"`
//	Zillya GoZillya `json:"Zillya"`
//	ZoneAlarm GoZoneAlarm `json:"ZoneAlarm"`
//	Zoner GoZoner `json:"Zoner"`
//	EGambit GoEGambit `json:"eGambit"`
//}
//
//type GoLastAnalysisStats struct {
//	ConfirmedTimeout int `json:"confirmed-timeout"`
//	Failure int `json:"failure"`
//	Harmless int `json:"harmless"`
//	Malicious int `json:"malicious"`
//	Suspicious int `json:"suspicious"`
//	Timeout int `json:"timeout"`
//	TypeUnsupported int `json:"type-unsupported"`
//	Undetected int `json:"undetected"`
//}
//
//type GoMainIcon struct {
//	Dhash string `json:"dhash"`
//	RawMd5 string `json:"raw_md5"`
//}
//
//type GoDocpropsApp struct {
//	AppVersion string `json:"AppVersion"`
//	Application string `json:"Application"`
//	Characters string `json:"Characters"`
//	CharactersWithSpaces string `json:"CharactersWithSpaces"`
//	DocSecurity string `json:"DocSecurity"`
//	HyperlinksChanged string `json:"HyperlinksChanged"`
//	Lines string `json:"Lines"`
//	LinksUpToDate string `json:"LinksUpToDate"`
//	Pages string `json:"Pages"`
//	Paragraphs string `json:"Paragraphs"`
//	ScaleCrop string `json:"ScaleCrop"`
//	SharedDoc string `json:"SharedDoc"`
//	Template string `json:"Template"`
//	TotalTime string `json:"TotalTime"`
//	Words string `json:"Words"`
//}
//
//type GoDocpropsCore struct {
//	CpLastModifiedBy string `json:"cp:lastModifiedBy"`
//	CpRevision string `json:"cp:revision"`
//	DcCreator string `json:"dc:creator"`
//	DctermsCreated time.Time `json:"dcterms:created"`
//	DctermsModified time.Time `json:"dcterms:modified"`
//}
//
//type GoMacros struct {
//	Length int `json:"length"`
//	Patterns []interface{} `json:"patterns"`
//	Properties []string `json:"properties"`
//	Sha256 string `json:"sha256"`
//	StreamPath string `json:"stream_path"`
//	Subfilename string `json:"subfilename"`
//	VbaCode string `json:"vba_code"`
//	VbaFilename string `json:"vba_filename"`
//}
//
//type GoOle struct {
//	Macros []GoMacros `json:"macros"`
//	NumMacros int `json:"num_macros"`
//}
//
//type GoLanguages struct {
//	ArSa int `json:"ar-sa"`
//	EnUs int `json:"en-us"`
//}
//
//type GoTypeContent struct {
//	Languages GoLanguages `json:"languages"`
//}
//
//type GoOpenxmlInfo struct {
//	ContentTypes []string `json:"content_types"`
//	DocpropsApp GoDocpropsApp `json:"docprops_app"`
//	DocpropsCore GoDocpropsCore `json:"docprops_core"`
//	FileType string `json:"file_type"`
//	Ole GoOle `json:"ole"`
//	Rels []string `json:"rels"`
//	TypeContent GoTypeContent `json:"type_content"`
//}
//
//type GoTotalVotes struct {
//	Harmless int `json:"harmless"`
//	Malicious int `json:"malicious"`
//}
//
//type GoTrid struct {
//	FileType string `json:"file_type"`
//	Probability float64 `json:"probability"`
//}
//
//type GoAttributes struct {
//	BundleInfo GoBundleInfo `json:"bundle_info"`
//	CapabilitiesTags []interface{} `json:"capabilities_tags"`
//	CreationDate int `json:"creation_date"`
//	Downloadable bool `json:"downloadable"`
//	Exiftool GoExiftool `json:"exiftool"`
//	FirstSubmissionDate int `json:"first_submission_date"`
//	LastAnalysisDate int `json:"last_analysis_date"`
//	LastAnalysisResults GoLastAnalysisResults `json:"last_analysis_results"`
//	LastAnalysisStats GoLastAnalysisStats `json:"last_analysis_stats"`
//	LastModificationDate int `json:"last_modification_date"`
//	LastSubmissionDate int `json:"last_submission_date"`
//	Magic string `json:"magic"`
//	MainIcon GoMainIcon `json:"main_icon"`
//	Md5 string `json:"md5"`
//	Names []interface{} `json:"names"`
//	OpenxmlInfo GoOpenxmlInfo `json:"openxml_info"`
//	Reputation int `json:"reputation"`
//	Sha1 string `json:"sha1"`
//	Sha256 string `json:"sha256"`
//	Size int `json:"size"`
//	Ssdeep string `json:"ssdeep"`
//	Tags []string `json:"tags"`
//	TimesSubmitted int `json:"times_submitted"`
//	Tlsh string `json:"tlsh"`
//	TotalVotes GoTotalVotes `json:"total_votes"`
//	Trid []GoTrid `json:"trid"`
//	TypeDescription string `json:"type_description"`
//	TypeExtension string `json:"type_extension"`
//	TypeTag string `json:"type_tag"`
//	UniqueSources int `json:"unique_sources"`
//	Vhash string `json:"vhash"`
//}
//
//type GoContextAttributes struct {
//	MatchInSubfile bool `json:"match_in_subfile"`
//	NotificationDate int `json:"notification_date"`
//	NotificationID string `json:"notification_id"`
//	NotificationSnippet string `json:"notification_snippet"`
//	NotificationSourceCountry interface{} `json:"notification_source_country"`
//	NotificationSourceKey interface{} `json:"notification_source_key"`
//	NotificationTags []string `json:"notification_tags"`
//	RuleName string `json:"rule_name"`
//	RuleTags []interface{} `json:"rule_tags"`
//	RulesetID string `json:"ruleset_id"`
//	RulesetName string `json:"ruleset_name"`
//}
//
//type GoLinks struct {
//	Self string `json:"self"`
//}
//
//type GoData struct {
//	Attributes GoAttributes `json:"attributes"`
//	ContextAttributes GoContextAttributes `json:"context_attributes"`
//	ID string `json:"id"`
//	Links GoLinks `json:"links"`
//	Type string `json:"type"`
//}
//
//type GoMeta struct {
//	Cursor string `json:"cursor"`
//}

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
	} `json:"data"`
	Meta struct {
		Cursor string `json:"cursor"`
	} `json:"meta"`
}

func LiveHunting(repo repository.IocRepo) {
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
					fmt.Println("sample->", sample)
				}
			}
		} else {
			cursor = cursor[:0]
		}
	}
	fmt.Println("len listSample->", len(sampleList))

	/*queue := helper.NewJobQueue(runtime.NumCPU())
	queue.Start()
	defer queue.Stop()
	for _, sample := range sample_list {
		queue.Submit(&LiveHuntingProcess{
			sample:  sample,
			iocRepo: repo,
		})
	}*/
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

type LiveHuntingProcess struct {
	sample  model.Sample
	iocRepo repository.IocRepo
}

func (process *LiveHuntingProcess) Process() {

}
