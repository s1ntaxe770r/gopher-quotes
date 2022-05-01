# Gopher Quotes 

Yet another application designed as a "MicRoserVice" for demo purposes. Gopher quotes offers users insghtful quotes ✍🏼 from their favourite gophers around the world. 🌍


## Architecture 🏗

Gopher Quotes is currently composed of three services, all of which communicate over HTTP.

![arch](./arch.png)


| Service    | Language/framework | description                                        |
|------------|--------------------|----------------------------------------------------|
| quotes-frontend| python/flask   | renders quotes from the api                        |
| quotes-api | Go                 | handles creation of quotes                         |
| quote-stats| python/flask       | provides "useful" stats on quotes                  |



## Setup using Docker-Compose  :whale:

Starting the service is as easy as:

```bash
docker-compose up
```

