# News Aggregator
 A microservices-based news aggregator application that fetches headline news from different Api endpoints
using Golang and the Gin framework. 

### Components:
- [News Aggregator Service]() Fetches news from various external APIs, aggregates the data, and standardizes the format.
- [User Service]() Manages user registration, authentication, and user preferences for news categories or sources.
- [Content Delivery Service]() Serves aggregated news to users based on their preferences, with functionalities like search and filtering. 


## Getting started
  The Application uses the [NewsApi](https://newsapi.org/) as its external Api accessing different api endpoints based on the request paramters made by the user. 
 ###
 Postman will be used in testing and trying out this application

### Signing Up 
   To Signup , Email, Name, and password fields are requeired as inputs
  ```
  {
    "name":"your_name",
    "email":"your_email",
    "password":"your_password"
 }
```
### Login 
 Login with your credentials used for signing up to get started with things
  ```
 {
    "email":"your_email",
    "password":"your_password"
 }
```
###  Search  
  ```
  {
    "country":"us",
    "category":"science"
 }
  ```
