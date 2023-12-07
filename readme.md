# News Aggregator
 A Backend microservices-based news aggregator application that fetches headline news from different Api endpoints
built with Golang and the Gin framework. 

## Tech Stacks Used
- [Golang](https://go.dev/) for the backend
- [Postgres](https://www.postgresql.org/) for relational database management for authentication and user credentials
- [Redis](https://redis.io/) for session management and caching
- [Google Cloud Platform](https://cloud.google.com/gcp?utm_source=google&utm_medium=cpc&utm_campaign=emea-ng-all-en-bkws-all-all-trial-e-gcp-1011340&utm_content=text-ad-none-any-DEV_c-CRE_501794636569-ADGP_Hybrid+%7C+BKWS+-+EXA+%7C+Txt+~+GCP+~+General%23v3-KWID_43700061569959215-aud-1651755615252:kwd-87853815-userloc_1010294&utm_term=KW_gcp-NET_g-PLAC_&&gad_source=1&gclid=CjwKCAiA98WrBhAYEiwA2WvhOhuA-uqXeFIV_xs_57NtnkVtmIvMO8Y8zjIb98uLiuRx4u6gRSzsGBoCxj4QAvD_BwE&gclsrc=aw.ds&hl=en) for deployment using the GCP compute engine 
  
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
 News artices mainly in headline formats are fetched based on country and category  inputs. The country and category inputs are used as search parameters.  These are the list of paramters that can be used to fetch the news formats
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
The outputs below gives the list of articles, it has various fields such as author, title, description which gives various information about a particular news item.
Please Note that you have to be authenticated i.e you must login to be able to fetch and access the api endpoints for you to receive news articles. 
#### Outputs 
 ```
{
    "message": "Preferences set successfully"
}{
    "Preferences set successfully": [
        {
            "source": {
                "id": "google-news",
                "name": "Google News"
            },
            "author": "Ars Technica",
            "title": "Daily Telescope: A super-hot jet 1000 light-years from Earth - Ars Technica",
            "description": "",
            "publishedAt": "2023-12-06T13:00:31Z",
            "content": ""
        },
        {
            "source": {
                "id": "",
                "name": "ScienceAlert"
            },
            "author": "Universe Today",
            "title": "China's Lander Detects Giant Polygonal Structures Buried Beneath Mars - ScienceAlert",
            "description": "China's Zhurong rover was equipped with a ground-penetrating radar system, allowing it to peer beneath Mars's surface.",
            "publishedAt": "2023-12-06T12:13:50Z",
            "content": "China's Zhurong rover was equipped with a ground-penetrating radar system, allowing it to peer beneath Mars's surface.\r\nResearchers have announced new results from the scans of Zhurong's landing site… [+4595 chars]"
        },
        {
            "source": {
                "id": "",
                "name": "Theregister.com"
            },
            "author": "Lindsay Clark",
            "title": "'Wobbly spacetime' is latest stab at unifying physics - The Register",
            "description": "Grudge match between quantum mechanics and general relativity attracts new effort to find harmony",
            "publishedAt": "2023-12-06T11:45:00Z",
            "content": "Since the early 20th century, physicists have struggled to marry theories governing the very big with those for the very small.\r\nDespite the staggering achievements in modern science, the conflict be… [+3009 chars]"
        },
        {
            "source": {
                "id": "",
                "name": "SpaceNews"
            },
            "author": "Jeff Foust",
            "title": "OSIRIS-REx parachute deployment affected by wiring error - SpaceNews",
            "description": "The drogue parachute on the OSIRIS-REx sample return capsule failed to deploy properly because of a design error that did not prevent a safe landing.",
            "publishedAt": "2023-12-06T09:04:41Z",
            "content": "WASHINGTON The drogue parachute on the OSIRIS-REx sample return capsule failed to deploy properly because of a design error, a flaw that did not prevent a safe landing of the capsule.\r\nIn a Dec. 5 st… [+3795 chars]"
        },
        {
            "source": {
                "id": "",
                "name": "IFLScience"
            },
            "author": "Dr. Alfredo Carpineti",
            "title": "Naturally Occurring Magnetic Monopoles Measured For The First Time - IFLScience",
            "description": "Quasiparticles have been found to only have a single magnetic pole, or two, or four.",
            "publishedAt": "2023-12-06T07:21:18Z",
            "content": "Regular magnets have two poles, a north and a south, and their behavior is defined in classical terms by the Maxwell equations. From contemporaries of Maxwell through to modern researchers, there hav… [+2456 chars]"
        },
        {
            "source": {
                "id": "",
                "name": "Space.com"
            },
            "author": "Robert Lea",
            "title": "Gravitational waves rippling from black hole merger could help test general relativity - Space.com",
            "description": "\"I never thought I would ever see such a measurement in my lifetime.\"",
            "publishedAt": "2023-12-06T01:00:18Z",
            "content": "Space is part of Future US Inc, an international media group and leading digital publisher. Visit our corporate site.\r\n©\r\nFuture US, Inc. Full 7th Floor, 130 West 42nd Street,\r\nNew York,\r\nNY 10036."
        },
        {
            "source": {
                "id": "",
                "name": "Space.com"
            },
            "author": "Monisha Ravisetti",
            "title": "NASA's Psyche spacecraft finds its 'first light' while zooming to a metal asteroid (image) - Space.com",
            "description": "The spacecraft's twin cameras have produced a 68-image stellar mosaic.",
            "publishedAt": "2023-12-05T22:38:50Z",
            "content": "Space is part of Future US Inc, an international media group and leading digital publisher. Visit our corporate site.\r\n©\r\nFuture US, Inc. Full 7th Floor, 130 West 42nd Street,\r\nNew York,\r\nNY 10036."
        },
        {
            "source": {
                "id": "",
                "name": "PetaPixel"
            },
            "author": "Jeremy Gray",
            "title": "How a Photographer Captured Aurora, STEVE, and the Milky Way in One Shot - PetaPixel",
            "description": "How a photographer captured an extremely rare night sky image.",
            "publishedAt": "2023-12-05T22:23:36Z",
            "content": "© NorthernPixl\r\nBritish photographer Stephen Pemberton, who goes by NorthernPixl, has demonstrated how determination, preparation, and good luck can result in beautiful night sky photographs. When ph… [+6942 chars]"
        },
        {
            "source": {
                "id": "",
                "name": "Phys.Org"
            },
            "author": "Science X",
            "title": "'Friendly' hyenas are more likely to form mobs, research shows - Phys.org",
            "description": "After more than 35 years of surveillance, Michigan State University researchers are exposing some of the secret workings of mobs. To be clear, these mobs are made up of spotted hyenas.",
            "publishedAt": "2023-12-05T20:06:04Z",
            "content": "After more than 35 years of surveillance, Michigan State University researchers are exposing some of the secret workings of mobs. To be clear, these mobs are made up of spotted hyenas.\r\nPublishing in… [+6698 chars]"
        }
    ]
}
 ```




