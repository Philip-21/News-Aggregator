# News Aggregator
 A microservices-based news aggregator application that fetches headline news from different Api endpoints
using Golang and the Gin framework. 

## Components:
- [News Aggregator Service](https://github.com/Philip-21/News-Aggregator/tree/master/newsaggregator-service) Fetches news from various external APIs, aggregates the data, and standardizes the format.
- [User Service](https://github.com/Philip-21/News-Aggregator/tree/master/user-service) Manages user registration, authentication, and user preferences for news categories or sources.
- [Content Delivery Service](https://github.com/Philip-21/News-Aggregator/tree/master/contentdelivery-service) Serves aggregated news to users based on their preferences, with functionalities like search and filtering. 


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
 Login with your credentials used in signing up to get started with things. 
  ```
 {
    "email":"your_email",
    "password":"your_password"
 }
```
###  Getting news content.  
 News artices mainly in headline formats are fetched based on inputs in the country and category field. These are the list of paramters that can be used to fetch the news formats
 #### Country 
   The country parameters to get the headline news are in 2 letter codes 
    `ae ar at au be bg br ca ch cn co cu cz de eg fr gb gr hk hu id ie il in it jp kr lt lv ma mx my ng nl no nz ph pl pt ro rs ru sa se sg si sk th tr tw ua us ve za`
 #### Category 
   The category parameters to get the headline news include
   `business, entertainment, generalhealth, science, sports, technology`
You  can look through the [api docs](https://newsapi.org/docs/endpoints/top-headlines) for more details on this 
 #### Inputs
  ```
  {
    "country":"us",
    "category":"science"
 }
  ```
#### Outputs 
 ```
 {
  "status": "ok",
  "totalResults": 10,
  "articles": [
    {
      "source": {
        "id": null,
        "name": "Big Think"
      },
      "author": "Ethan Siegel",
      "title": "\"Singularities don't exist,\" claims black hole pioneer Roy Kerr - Big Think",
      "description": "The brilliant mind who discovered the spacetime solution for rotating black holes claims singularities don't physically exist. Is he right?",
      "url": "https://bigthink.com/starts-with-a-bang/singularities-dont-exist-roy-kerr/",
      "urlToImage": "https://bigthink.com/wp-content/uploads/2023/12/hikerr-negview.png?w=1024&h=576&crop=1",
      "publishedAt": "2023-12-05T07:00:00Z",
      "content": null
    },
    {
      "source": {
        "id": null,
        "name": "ScienceAlert"
      },
      "author": "Michelle Starr",
      "title": "We Might Be Sitting in a Massive 'Supervoid' in Space, And That Could Explain The Hubble Tension - ScienceAlert",
      "description": "When we gaze out into the cosmos beyond the borders of the Milky Way, we behold multitudes.",
      "url": "https://www.sciencealert.com/we-might-be-sitting-in-a-massive-supervoid-in-space-and-that-could-explain-the-hubble-tension",
      "urlToImage": "https://www.sciencealert.com/images/2023/12/Kroupa-Hubble-Spannung.jpg",
      "publishedAt": "2023-12-05T05:53:33Z",
      "content": "When we gaze out into the cosmos beyond the borders of the Milky Way, we behold multitudes. Space is teeming with galaxies, speckled across the darkness like stars. If we stopped there, it would be e… [+4380 chars]"
    },
   }
  ]
}
 ```
The outputs above gives the list of articles, it has various fields such as author, title, description which gives various information about a particular news item
Please Note that you have to be authenticated to be able to fetch and access the api endpoints for you to receive news articles 

