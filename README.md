# golang_k8s_cronjob

a simple golang app that makes an Http Request to the TradeMade public api and creates a k8s cronjob to execute the request every 2 minute

# instalation with minikube and kubernetes

- make sure you have minikube and kubernetes already installed

- execute the following steps

- create the docker image for the cronjob with: docker build -t trade-made-cron -f Dockerfile .

- generate all the kubernetes resources running: kubectl apply -f k8s/

- the resources created by the previous command generate the following k8s resources: 

- 1 cronjob with 1 POD.
- 1 secret for passing enviornment variables like the Api Key for the TradeMade public API. 
