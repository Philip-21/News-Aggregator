# News Aggregator
 A Backend microservices-based news aggregator application that fetches headline news from different Api endpoints
using Golang and the Gin framework. 

## Tech Stacks Used
- Golang
- Postgres
- Redis
- Google Cloud Platform
  
## Components:
- [News Aggregator Service](https://github.com/Philip-21/News-Aggregator/tree/master/newsaggregator-service) Fetches news from various external APIs, aggregates the data, and standardizes the format.
- [User Service](https://github.com/Philip-21/News-Aggregator/tree/master/user-service) Manages user registration, authentication, and user preferences for news categories or sources.
- [Content Delivery Service](https://github.com/Philip-21/News-Aggregator/tree/master/contentdelivery-service) Serves aggregated news to users based on their preferences, with functionalities like search and filtering. 


## Getting started
  The Application uses the [NewsApi](https://newsapi.org/) as its external Api accessing different api endpoints based on the request paramters made by the user. 
 ###
 [Postman](https://www.postman.com/) will be used in testing and trying out this application

### Viewing the home page to ensure you are connected
  - URL http://34.171.120.144/
  - Method = GET
   ```
   {
    "message": "welcome to news headline service"
   }
   ```
### Signing Up 
  - To Signup , Email, Name, and password fields are requeired as inputs
  - URL http://34.171.120.144/user/signup
  - Method = POST
  ```
  
 {
    "name":"your_name",
    "email":"your_email",
    "password":"your_password"
 }
```
### Login 
 - Login with your credentials used in signing up to get started with things.
 - URL http://34.171.120.144/user/login
 - Method = POST
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
- URL http://34.171.120.144/news/user/preference
- Method = POST
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
     {
       "source": {
         "id": null,
         "name": "SpaceNews"
       },
       "author": "Jeff Foust",
       "title": "India returns Chandrayaan-3 propulsion module to Earth orbit - SpaceNews",
       "description": "The spacecraft that transported the Chandrayaan-3 lander to the moon has returned to Earth orbit, testing technologies for a lunar sample return mission.",
       "url": "https://spacenews.com/india-returns-chandrayaan-3-propulsion-module-to-earth-orbit/",
       "urlToImage": "https://spacenews.com/wp-content/uploads/2023/12/cyyan3-propmodule.jpg",
       "publishedAt": "2023-12-05T03:51:19Z",
       "content": "WASHINGTON The spacecraft that transported the Chandrayaan-3 lander to the moon has returned to Earth orbit, a demonstration of technologies to support a future Indian lunar sample return mission.\r\nT… [+3943 chars]"
     },
     {
       "source": {
         "id": null,
         "name": "Space.com"
       },
       "author": "Sharmila Kuthunur",
       "title": "Pluto's 'almost twin' dwarf planet Eris is surprisingly squishy - Space.com",
       "description": "New models show Eris is behaving less like a solid, rocky world and more like \"soft cheese.\"",
       "url": "https://www.space.com/pluto-twin-dwarf-planet-eris-squishy",
       "urlToImage": "https://cdn.mos.cms.futurecdn.net/5wA9HMEuB49dhKPdfv3Vd9-1200-80.jpg",
       "publishedAt": "2023-12-05T01:00:01Z",
       "content": "Space is part of Future US Inc, an international media group and leading digital publisher. Visit our corporate site.\r\n©\r\nFuture US, Inc. Full 7th Floor, 130 West 42nd Street,\r\nNew York,\r\nNY 10036."
     },
     {
       "source": {
         "id": null,
         "name": "PBS"
       },
       "author": null,
       "title": "Solar system with 6 planets orbiting in-sync discovered - PBS NewsHour",
       "description": "Astronomers have discovered a rare solar system with six planets moving in sync with one another. Estimated to be billions of years old, the formation 100 light-years away may help unravel some mysteries of our solar system. Miles O’Brien reports on this perf…",
       "url": "https://www.pbs.org/newshour/show/solar-system-with-6-planets-orbiting-in-sync-discovered-in-milky-way",
       "urlToImage": "https://d3i6fh83elv35t.cloudfront.net/static/2023/12/space-1024x683.jpg",
       "publishedAt": "2023-12-04T23:15:29Z",
       "content": "Miles OBrien:\r\nWell, there's two instruments involved.\r\nThere was a NASA instrument called the Transiting Exoplanet Survey Satellite, or TESS. TESS has been in orbit for about five years, and it has … [+543 chars]"
     },
    
   }
  ]
}
 ```
The outputs above gives the list of articles, it has various fields such as author, title, description which gives various information about a particular news item.
Please Note that you have to be authenticated i.e you must signup to be able to fetch and access the api endpoints for you to receive news articles. 



