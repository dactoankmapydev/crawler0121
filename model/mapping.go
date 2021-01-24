package model

const (
	IndexNameSample       = "sample"
	IndexNameIndicator       = "indicator"
	IndexNamePost       = "post"
	IndexNameCompromised       = "compromised"

	MappingSample = `
	{
		"settings":{
			"number_of_shards": 2,
			"number_of_replicas": 0
		},
		"mappings":{
			"properties":{
				"names":{
					"type":"text"
				},
				"sha256":{
					"type":"text"
				},
				"sha1":{
					"type":"text"
				},
				"md5":{
					"type":"text"
				},
				"first_submit":{
					"type":"date"
				},
				"notification_date":{
					"type":"date"
				},
				"file_type":{
					"type":"text"
				},
				"tags":{
					"type":"text"
				},
				"engines_detected":{
					"type":"text"
				},
				"detected":{
					"type":"integer"
				},
				"point":{
					"type":"integer"
				}
			}
		}
	}`

	MappingIndicators = `
	{
		"settings":{
			"number_of_shards": 2,
			"number_of_replicas": 0,
		},
		"mappings":{
			"properties":{
				"ioc_id":{
					"type":"text"
				},
				"ioc":{
					"type":"text"
				},
				"ioc_type":{
					"type":"text"
				},
				"crawled_time":{
					"type":"text"
				},
				"source":{
					"type":"text"
				},
				"category":{
					"type":"text"
				}
			}
		}
	}`

	MappingPost = `
	{
		"settings":{
			"number_of_shards": 2,
            "number_of_replicas": 0
		},
		"mappings":{
			"properties":{
				"id":{
					"type":"text"
				},
				"name":{
					"type":"text"
				},
				"description":{
					"type":"text"
				},
				"author_name":{
					"type":"text"
				},
				"tags":{
					"type":"text"
				},
				"targeted_countries":{
					"type":"text"
				},
				"malware_families":{
					"type":"text"
				},
				"attack_ids":{
					"type":"text"
				},
				"references":{
					"type":"text"
				},
                "industries":{
					"type":"text"
				}
			}
		}
	}`

	MappingCompromised = `
	{
		"settings":{
			"number_of_shards": 2,
			"number_of_replicas": 0
		},
		"mappings":{
			"properties":{
				"uid":{
					"type":"text"
				},
				"hostname":{
					"type":"text"
				},
				"src":{
					"type":"text"
				},
				"victim_hash":{
					"type":"text"
				},
				"creation_date":{
					"type":"date"
				},
				"timestamp":{
					"type":"date"
				},
				"country":{
					"type":"text"
				}
			}
		}
	}`
)
