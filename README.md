# cmonitor
Lists/Subscribers implementation with a Dockerized golang app (https://www.campaignmonitor.com)

The app is a simple implementation that allows a user to view their mailing lists from Campaign Monitor service and add or remove subscribers from each list.

The server side component handling the requests is written in Go using only the native packages. The client component is written in javascript using Vue js (https://vuejs.org/) and bulma css framework (https://bulma.io/). 

### Usage
Before using the app you need to get the API Credentials from campaign monitor and set them up as environment variables:
1. CM_API_KEY
2. CM_API_CLIENT

The app exposes port 8080 so make sure that no other application is using that port.

Select the way to run the app that suits you the most:

* Using the Dockerfile to build and run the container:
```
docker build -t supergramm/cmonitor .
docker run -d -p 8080:8080 --name cmonitor -e CM_API_KEY=you_api_key -e CM_API_CLIENT=your_api_client_key supergramm/cmonitor
```


* Using Docker Compose tool:
```
docker-compose up -d
```


* Using go (static folder must sit on the same level as the executable to serve static files):
```
go get github.com/speix/cmonitor
go install github.com/speix/cmonitor
```


### Docker image hosted on DockerHub
https://hub.docker.com/r/supergramm/cmonitor/

### Preview
<p float="left">
<img src="http://www.supergramm.com/media/images/github/cmonitor2.jpg" alt="" data-canonical-src="http://www.supergramm.com/media/images/github/cmonitor1.jpg" height="500"/>
<img src="http://www.supergramm.com/media/images/github/cmonitor1.jpg" alt="" data-canonical-src="http://www.supergramm.com/media/images/github/cmonitor1.jpg" height="500"/>
</p>
