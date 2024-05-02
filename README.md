# golang_gin_chatroom
1. zsh cmd: docker build -t gochatroom:latest .
   // build image
2. minikube ref: https://minikube.sigs.k8s.io/docs/start/
   // 參考 所使用的 k8s core etc. minikube, k8s, k3d, colima
3. zsh cmd: minikube start
4. into minikube : kubectl apply -f deploy/deployment.yaml
5. into minikube : kubectl apply -f deploy/service.yaml
   // now deplo done
6. minikube service 
6. zsh cmd: kubectl get svc -A get service nodeport to visit
   // for me environment : http://192.168.106.4:32161/ 32161 is svc node port
