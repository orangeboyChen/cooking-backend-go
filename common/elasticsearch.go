package common

import (
	"context"
	"cooking-backend-go/entity"
	"github.com/olivere/elastic/v7"
	"log"
	"time"
)

var ESClient *elastic.Client
var CourseIndex = "cooking_course"
var TagIndex = "cooking_tag"

func InitElasticSearch() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	ESClient = client

	//_, err = client.DeleteIndex(CourseIndex).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	prepareCourseIndex()
	prepareTagIndex()

	var b = entity.SearchCourse{
		Name:       "中国国家",
		Detail:     "1",
		Image:      "1",
		TagId:      "1",
		UserId:     "1",
		UserAvatar: "1",
		CreateTime: time.Now().UnixMilli(),
	}

	res, err := ESClient.Index().Index(CourseIndex).BodyJson(&b).Do(context.Background())
	if err != nil {
		log.Panic(err)
	}

	log.Print(res)
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
			log.Fatal("创建cooking_course索引失败", err.Error())
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
			log.Fatal("创建cooking_tag索引失败", err.Error())
		}
	}
}
