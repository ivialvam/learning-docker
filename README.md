# learning-docker

step by step

1. Run Existing Docker

2. Open cmd, run a new Container
- docker run -d -p 8080:80 --name welcome1 docker/welcome1-to-docker
  ![-f --tail 10 welcome1](https://github.com/ivialvam/learning-docker/assets/97967090/bf0d9966-a824-4a26-a73e-4ad47f1e9676)

3. Coba dibrowser untuk mengakses apps didalam welcome-to docker image
- buka URL http://localhost:8080
![welcome11](https://github.com/ivialvam/learning-docker/assets/97967090/a6cb6147-2ff5-4fd5-ba33-61e60557eddc)

4. Execute sebuah perintah saat container welcome1 sedang berjalan

5. Stop Container dan Melihat log 10 baris proses terakhir yang berjalan di Container welcome1
![-f --tail 10 welcome1](https://github.com/ivialvam/learning-docker/assets/97967090/bc3adc0d-32a1-4326-a995-bf3e9e0a8d5b)


# Proses dan Output Task: Build Docker to Docker Image

![dockerbuild](https://github.com/ivialvam/learning-docker/assets/97967090/c1f50f47-601c-4299-8162-3ae36a1969ba)
![docker build to docker image](https://github.com/ivialvam/learning-docker/assets/97967090/62d84fbc-69a5-4d8a-a3f8-5392ff737e71)
![output](https://github.com/ivialvam/learning-docker/assets/97967090/e41867dc-30f4-40ad-8943-e338918eaa4a)
![dc log](https://github.com/ivialvam/learning-docker/assets/97967090/1b480bb8-f782-47c3-81af-eaca3b9a46d5)

