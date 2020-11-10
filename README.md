# ATM Withdrawal Analysis
Based in real coins with value 50, 10, 5, and 1, this api calculate how many coins will be used based on value.

##Install
Requirements:
* docker 19 or latest
* docker-compose 1.27 or latest
* make command

#### Commands:

##### Unit test:

```shell script 
$ make test 
```

##### Run application:
```shell script
$ make run-docker \\ Runs docker-compose with application image and its requirements
```

```shell script
$ make run-docke-clean \\ Run a clean build image from docker application 
```

## Project Organization
### Third Part Packages
* [github.com/go-redis/redis/v8](https://github.com/go-redis/redis) - Main redis package from golang community
* [github.com/google/uuid](https://github.com/google/uuid) - Main package for uses uuid
* [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - Easy way to log the application
* [github.com/stretchr/testify](https://github.com/stretchr/testify) - Used for unit test assertions and mocks
### Withdrawal Flow
The application will consult the redis before to analise the value.

<img src="http://www.plantuml.com/plantuml/png/fP91RnD138NlyojyuC8XR5fnYgAgeG8gLK2LH70eFRYx3XbfToQCFOa2uh-pa-oahfQII9sZPN_Fy_mkMJ18R6eZLcg2CTjVD3eVNi0tAyigu0PhkA8griGmSLXRtlAuI5qNPk6zK02OOgojy_2Pn7zt07TtIv4LZuhYHmk93szAtnHXt-H8V87I33O7_W0xBapcdxHkrhk_DyPWAJv0P0fcXK5iQLVeHBKRSc9b2zVp-UsamfLGoAKPXRKjoLGe-81CCANAiaam6EYlZk3Z-paKwxFYSRtYrMziV7Jo9-rWroYaLTM-BKg9Z-3iDTm6Nn9pAa_fPoGM0-qxWFIoq3rLwo6_OcshDAXfEs8jO1Uv39R1kdLgpXtQsF5Wy1Gx9Od1vP4UV6TK6j9E9G3C12Kol1qEuHts1ybb_Vx7ybL6C6lRG7ms2oR9xWnHPdK3YfXB2RUxQUiG1RlOmXMMIqeBSQo5yuZ2vH3qOUcdcTOPymzCHS-2yZ6kaVThxCTLHy-E99dbioftil31H-tZqPIq0w51IUChxrWRTc8Ib7GY2Mdvqo9ntaRYouqQqs7GRkiav0IKtZfk4yLKWNLT85Bu8vdZ3Xv-aNc78Fm_EZCIk75-a-yDulSvEXwV_FRL0nDQAAxf2VMa45NKJ_sBZz-x_7D6zGwlrb5QvjGGF8whJdD9hefD_HS0">

### Directory Structure
```
   ├── build ##Docker configuration
   │   └── docker
   │       ├── Dockerfile
   │       └── start.server.sh
   ├── config
   │   ├── container ## Injection Container
   │   │   ├── container.go
   │   │   └── container_test.go
   │   └── routes ## Routes
   │       ├── routes.go
   │       └── routes_test.go
   ├── docker-compose.yml
   ├── docs 
   │   ├── swagger.yml
   │   └── withdrawal.get.v1.plantuml
   ├── go.mod
   ├── go.sum
   ├── internal
   │   ├── httpserver ## net/http abstraction
   │   │   ├── errors.go
   │   │   ├── httpserver.go
   │   │   ├── httpserver_test.go
   │   │   └── responses.go
   │   ├── primary ## Primari ports
   │   │   └── v1.withdrawalhttp
   │   │       ├── withdrawal.go
   │   │       └── withdrawal_test.go
   │   └── secondary ## Secondary ports
   │       └── cache
   │           ├── redis.go
   │           └── redis_test.go
   ├── LICENSE
   ├── main.go
   ├── Makefile
   ├── pkg
   │   └── domain ## main logic
   │       └── v1.withdrawal
   │           ├── enum.go
   │           ├── service.go
   │           ├── service_test.go
   │           └── types.go
   ├── README.md
   └── tools
       └── logger
           ├── logger.go
           └── logger_test.go
```
## Biography

1. [My Talk About Hexagonal Architecture](https://docs.google.com/presentation/d/1nEpfDEfnwGB3Xy-7CMfccW7L2qZeo4UVR738434bVVY/edit?usp=sharing)
2. [My Talk About Postman](https://docs.google.com/presentation/d/1SHUSATWs-vOkScWXm6ae4vgjomEKeDrMkm0JQ1fXpRw/edit?usp=sharing)
3. [Clean Architecture - Uncle Bob](https://www.amazon.com.br/Clean-Architecture-Craftsmans-Software-Structure-ebook/dp/B075LRM681/)
4. [The Proposed Project Structure](https://github.com/golang-standards/project-layout)

## License

 [MIT license](https://github.com/raulinoneto/atm-withdrawal-analisys/blob/main/LICENSE)


