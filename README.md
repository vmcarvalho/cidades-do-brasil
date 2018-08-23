# Cidades do Brasil

Golang script to generate a database containing all brazilian cities, including geolocation.
- Source Data is fetched from government official repository CSV file. 
- Stores in Mongodb database. Easily extandable to other formats!

## Use

```
$ docker-compose up
```

Example:
```
$ docker-compose up -d
go_1       | Starting Go Parser
go_1       | ... parsing input
go_1       | ... dumping to database  mongodb://mongodb:27017/CidadesDoBrasil
go_1       | ... Finished!
cidades-do-brasil_go_1 exited with code 0
$ mongo CidadesDoBrasil --eval "printjson(db.City.findOne())"
MongoDB shell version: 2.6.10
connecting to: CidadesDoBrasil
{
	"_id" : ObjectId("5b6fb77c6a52e6ffbf4c3721"),
	"name" : "Abadia de Goiás",
	"uf" : "GO",
	"lat" : -16.7855,
	"lon" : -49.4538
}
```