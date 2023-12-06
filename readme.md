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
     {
       "source": {
         "id": null,
         "name": "Space.com"
       },
       "author": "Robert Lea",
       "title": "How Perseverance rover captured its youngest samples from Mars' Dream Lake (video) - Space.com",
       "description": "\"Comparing rocks from different ages is important because Mars, like Earth, had environments that evolved and changed through time.\"",
       "url": "https://www.space.com/perseverance-mars-rover-youngest-material-yet-video",
       "urlToImage": "https://cdn.mos.cms.futurecdn.net/cULLGpQEZ8z6UksJE5dHd9-1200-80.gif",
       "publishedAt": "2023-12-04T22:32:18Z",
       "content": "Space is part of Future US Inc, an international media group and leading digital publisher. Visit our corporate site.\r\n©\r\nFuture US, Inc. Full 7th Floor, 130 West 42nd Street,\r\nNew York,\r\nNY 10036."
     },
     {
       "source": {
         "id": "cnn",
         "name": "CNN"
       },
       "author": "Allison Chinchar",
       "title": "Modern earthquakes in US could be aftershocks from quakes in the 1800s, scientists say - CNN",
       "description": "Aftershocks from devastating earthquakes in the 1800s near the Missouri-Kentucky border and in Charleston, South Carolina, may still be occurring, a study found.",
       "url": "https://www.cnn.com/2023/12/04/world/earthquake-century-old-aftershocks-us-scn/index.html",
       "urlToImage": "https://media.cnn.com/api/v1/images/stellar/prod/hjk00013.jpg?c=16x9&q=w_800,c_fill",
       "publishedAt": "2023-12-04T21:32:00Z",
       "content": "Sign up for CNNs Wonder Theory science newsletter. Explore the universe with news on fascinating discoveries, scientific advancements and more.\r\nAfter large earthquakes, there is an expectation that … [+8040 chars]"
     },
     {
       "source": {
         "id": null,
         "name": "Gothamist"
       },
       "author": "https://gothamist.com/staff/rosemary-misdary",
       "title": "Geminid meteor shower will be bright enough to be seen in NYC this month - Gothamist",
       "description": "December astronomy highlights include the winter solstice, the shortest night of the year, and one of the greatest meteor showers.",
       "url": "https://gothamist.com/news/geminid-meteor-shower-will-be-bright-enough-to-be-seen-in-nyc-this-month",
       "urlToImage": "https://cms.prod.nypr.digital/images/342160/fill-1200x650|format-webp|webpquality-85/",
       "publishedAt": "2023-12-04T21:32:00Z",
       "content": "This month, New Yorkers have a rare opportunity to enjoy a meteor shower in city parks or from the rooftops of apartment buildings.\r\nThe Geminids meteor shower promises around 120 shooting stars per … [+4599 chars]"
     },
     {
       "source": {
         "id": null,
         "name": "Teslarati"
       },
       "author": "Richard Angle",
       "title": "SpaceX continues its march to 100 launches in a year - TESLARATI",
       "description": "SpaceX is nearing 100 launches for 2023 as the year comes closer to its end.",
       "url": "https://www.teslarati.com/spacex-100-launches-2023/",
       "urlToImage": "https://www.teslarati.com/wp-content/uploads/2023/12/GAZjdfAbEAAfQpE-scaled.jpg",
       "publishedAt": "2023-12-04T21:07:41Z",
       "content": "Late Saturday night, a Falcon 9 lifted off from Space Launch Complex 40 at Cape Canaveral Space Force Station with another 23 Starlink satellites delivered into orbit.\r\nLaunch occurred at 11:00 pm ET… [+2228 chars]"
     },
   }
  ]
}
 ```
The outputs above gives the list of articles, it has various fields such as author, title, description which gives various information about a particular news item
Please Note that you have to be authenticated to be able to fetch and access the api endpoints for you to receive news articles 

