package elastic_config

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

var ESClient *elastic.Client
var CourseIndex = "cooking_course"
var TagIndex = "cooking_tag"

func InitElasticSearch() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Panic("你妈的没准备好数据库就来找我？你把我当谁了？艹你妈", err)
	}

	ESClient = client
	prepareCourseIndex()
	prepareTagIndex()
}

func prepareCourseIndex() {
	exists, _ := ESClient.IndexExists(CourseIndex).Do(context.Background())
	if !exists {
		log.Print("开始创建索引cooking_course")

		mapping := `{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "name": {
        "type": "text",
        "analyzer": "ik_smart",
        "search_analyzer": "ik_smart"
      },
      "detail": {
        "type": "text",
        "analyzer": "ik_smart",
        "search_analyzer": "ik_smart"
      },
      "image": {
        "type": "keyword"
      },
      "tagId": {
        "type": "keyword"
      },
      "userId": {
        "type": "keyword"
      },
      "userAvatar": {
        "type": "keyword"
      },
      "createTime": {
        "type": "date"
      }
    }
  }
}`

		_, err := ESClient.CreateIndex(CourseIndex).Body(mapping).Do(context.Background())
		if err != nil {
			log.Fatal("妈的创建cooking_course索引失败了，你自己去手动创建一个", err.Error())
		}
	}
}

func prepareTagIndex() {
	exists, _ := ESClient.IndexExists(TagIndex).Do(context.Background())
	if !exists {
		log.Print("开始创建索引cooking_tag")

		mapping := `{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "name": {
        "type": "text",
        "analyzer": "ik_smart",
        "search_analyzer": "ik_smart"
      },
      "tagTypeId": {
		"type": "keyword"
      }
    }
  }
}`

		_, err := ESClient.CreateIndex(TagIndex).Body(mapping).Do(context.Background())
		if err != nil {
			log.Fatal("妈的创建cooking_tag索引失败了，你他妈自己去手动创建一个", err.Error())
		}
	}
}
