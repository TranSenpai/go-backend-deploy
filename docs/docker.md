<!-- build image -->
docker build -t img-go-backend:latest . (podman build) -> Build ra image, cờ -t(tag) là tên image dấu . là đường dẫn Dockerfile

<!-- Xóa image -->
docker image rm img-go-backend (hoặc podman image rm / podman rmi)

<!-- Xóa cache -->
docker builder prune -af (với podman: podman system prune -f)

<!-- dive trên github là thư viện dùng để bóc tách 1 cái image ra bên trong nó có cái gì -->
# Cách 1: Xuất image ra tar rồi chạy dive (nếu cài dive qua Snap)
podman save localhost/img-go-backend:latest -o img.tar
dive docker-archive://img.tar

# Cách 2: Chạy trực tiếp (yêu cầu cài dive bằng file .deb/binary thay vì Snap)
dive img-go-backend --source podman
# hoặc
dive podman://img-go-backend

<!-- List các image -->
docker image ls

<!-- Chạy image, start container và run -->
docker run --name con-go-backend -d -p 8000:8080 --env-file .env img-go-backend:latest
podman run --name con-go-backend -d -p 8000:8080 --env-file .env.production img-go-backend:latest
<!-- Coi terminal container log real time -->
docker logs -f con-go-backend

<!-- Xóa container -->
docker container rm con-go-backend

<!-- Stop container -->
docker container stop con-go-backend

<!-- Start container -->
docker container start con-go-backend

<!-- Xóa container -->
docker container restart con-go-backend


<!-- List container  -->
docker container ls

docker network create go-network

docker compose --env-file .env.production up -d
# hoặc với podman-compose:
podman-compose --env-file .env.production up -d
