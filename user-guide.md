User Guide
==========
* First of all you need to run the main program to initiate the API, Controller and Scheduler

* First you have to login
curl -u username:password http://localhost:8080/login
* Then you need to create a workload
curl -d 'workload-id=filtro1&filter=grayscale' -H "Authorization: " http://localhost:8080/workloads
* Now you can upload an image that is going to be processed
curl -F 'data=@/home/dev/Downloads/download.jpeg' -H "Authorization: " http://localhost:8080/images
* Returns an status
curl -H "Authorization: " http://localhost:8080/status


* Running new worker
cd worker/
export GO111MODULE=off
go run main.go --controller <host>:<port> --worker-name <worker_name> --tags <tag1>,<tag2>
