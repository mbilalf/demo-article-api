# DEMO Article Restful JSON API
A GO implementation.

## Running 
To Setup Mysql run following docker commands. 

```
docker image pull mysql:5.7
docker container run -d -p 3306:3306 --name=mysql_local -e MYSQL_ROOT_PASSWORD=fidodido -e MYSQL_DATABASE=demo_db mysql:5.7
```

Run API Service
```sh
go run *.go
```

Server runs at localhost:8000

## Endpoints

### Get All Articles
GET /api/v1/articles

### Get Article by Id
GET /api/v1/articles/{id}

### Add new Articles
POST /api/v1/articles

Sample Request Body:
```
{
	"title": "In to the Wild",
	"date": "2018-08-12T15:15:18+10:00",
	"body": "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. ",
	"tags": ["adventure", "wild"]
}
```

### Search Articles
GET /api/v1/tags/{tag}/{date}

e.g. localhost:8000/api/v1/tags/health/2018-08-12

Response:
```
{
    "Tag": "health",
    "Count": 2,
    "Articles": [
        "1",
        "2"
    ],
    "RelatedTags": [
        "adventure",
        "fun"
    ]
}
```
