# DEMO Article Restful JSON API
A GO implementation

## Running 
```sh
go run *.go
```

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
