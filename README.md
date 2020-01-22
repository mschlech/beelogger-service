# beelogger-service
beelogger-service

![pipeline status](https://gitlab.com/mschlech/beelogger-service/badges/master/pipeline.svg)](https://gitlab.com/mschlech/beelogger-service/commits/master)

# create a mongodb container 
docker run --name mongohive -v $(pwd):/data/db -p 27017:27017 -d mongo:latest