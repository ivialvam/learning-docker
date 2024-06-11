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

# Run postgres with existing volume
sh
docker run -d |
-name postgres l
-e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pasSword |
-v my-postgres-data:/var/lib/postgresql/data \
-p 5431:5432 postgres

1. Menjalankan perintah docker run
![postgress dwoload](https://github.com/ivialvam/learning-docker/assets/97967090/cbde2e2a-d417-40c1-806e-f13bec959861)
> File sudah tersedia di Container
![docker postgres running](https://github.com/ivialvam/learning-docker/assets/97967090/635d9a13-ba33-4605-bf0f-6e77d6278004)
> dan Volume
![docker volume](https://github.com/ivialvam/learning-docker/assets/97967090/38df6629-a687-4591-95fa-49e2147ff7a4)

2. Membuat table
![db table](https://github.com/ivialvam/learning-docker/assets/97967090/72cb06e9-41c0-49e9-85f9-920deee64822)
> Tampilan Container di Docker
![docker postgres stopped](https://github.com/ivialvam/learning-docker/assets/97967090/1c9535a1-0662-4be7-b9a8-83d8c547f4fc)
![dockerpostgres](https://github.com/ivialvam/learning-docker/assets/97967090/aebcaee8-cc86-495c-839b-421a6a58e36e)

3. Membuat volume dengan nama server postgres, username ivialva, password 123456, port 5432:5432
![postgres v2](https://github.com/ivialvam/learning-docker/assets/97967090/b3305eae-13ad-4009-b940-5832db844c48)

> Tampilan output dari Continer yang sudah tersedia di Docker
![docker running container](https://github.com/ivialvam/learning-docker/assets/97967090/a6978e52-05f0-4f8f-b4a4-16e257b27e47)
> Tampilan output dari Image dengan nama postgres yang sudah dibuat
![docker running image](https://github.com/ivialvam/learning-docker/assets/97967090/94a8c2c8-dcc0-4ca2-ab3d-5eafd8c33d8a)
> Tampilan output dari Volume dengan nama my-postgres-v2-ivialvam
![volume v2](https://github.com/ivialvam/learning-docker/assets/97967090/63b93084-b6a2-47ef-9314-8a56e465be71)

4. Cek tabel yang sudah dibuat (masih ada)
![db table](https://github.com/ivialvam/learning-docker/assets/97967090/67314a3e-e834-400e-b2e4-779dff15519b)
