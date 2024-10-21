# Microservice_3service
A microservice system consisting 3 services: ``public-layer-service``,``user-service`` and ``listing-service``. Build using ``Golang version 1.23.2``,``GORM`` as ORM,``PostgreSQL`` as databases and ``Docker`` for containerization.
# High Level Overview
![img](https://drive.google.com/uc?export=view&id=1-bLq2BL6tLA0KJL7Hf3zGi81_23jdjyu)  
# Prerequisite
1. Have ``Golang`` min ``version 1.23.2`` or higher installed on your device
2. Have ``Docker`` installed on your device
# Usage
1. clone this repository and cd to ``Microservice_3service``
```
git clone https://github.com/W-ptra/Microservice_3service.git
cd Microservice_3service
```
2. run following script to activate the ``docker-compose.yaml``
```
docker-compose up
```
3. Use browser/postman/curl to interact with api
```
GET http://127.0.0.1:8000/public-api/listing
or
GET http://127.0.0.1:8000/public-api/listing?userId=1&pageNum=2&pageSize=5
Parameters:
    pageNum = int # Default = 1
    pageSize = int # Default = 10
    userId = str # Optional

POST http://127.0.0.1:8000/public-api/listing
headers:    
    Content-Type: application/json
Request body: (JSON body)
    { // example
        "userId": 1,
        "listingType": "rent", // either "rent" or "sale"
        "price": 6000 // can't negative
    }

POST http://127.0.0.1:8000/public-api/users
headers:    
    Content-Type: application/json
Request body: (JSON body)
    { // example
        "name": "Lorel Ipsum"
    }
```
4. To clean up run following script
```
docker-compose down
```
