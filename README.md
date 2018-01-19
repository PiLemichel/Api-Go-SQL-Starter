# Api-Go-SQL-Starter

This project aims to accelerate the development of your api during a hackathon or other.

## Getting Started

### Prerequisites

For launch the sever you need Golang, if you don't have early install go to:
```
https://golang.org/doc/install
```

### Installing


get the project on your local machine.
```
go get github.com/PiLemichel/Api-Go-SQL-Starter/
```
The project use five more library so you need download too.
```
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm/dialects/mysql
go get net/http
```

Change the path of connetion to your database:
```
 db, err = gorm.Open("mysql", "root:root@tcp(localhost:8889)/user") => //"username:password@tcp(adresse:port)/dbname
```

Now
```
go build
```
And enjoy ! :)


## Authors

* **Pierre Lemichel** - *Initial work* - [PiLemichel](https://github.com/PiLemichel)


## License

This project is licensed under the MIT License
