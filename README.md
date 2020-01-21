# beelogger-service
beelogger-service


# create a mongodb container 
docker run --name mongohive -v $(pwd):/data/db -p 27017:27017 -d mongo:latest