# DEMO Article Restful JSON API
A GO implementation.

## TODOS
- DB connection is created before every db call. Need to move it at a shared place. Need to work on shared context.
- TX handling is not tested properly, not sure if it working.
- Exeption handling and logging mechanism needs improvement.

- Valid http status codes are not being returned yet. Handle error by sending appropriate HTTP error code
- MVC layering needs improvement. View classes (questionable package name) are meant for http response. Repository

- Create articles call dont return saved data.
- /api/v1/articles Accepts only 1 article for now.
- /api/v1/article/{id} return 501 as repositry function to load arricle is not implemented



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

### Initlize Database
GET /api/v1/setup
Setup Database and add dummy data

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
