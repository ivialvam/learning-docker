# Task: Run Existing Docker Image as a Container

1. Run command in terminal docker run -d -p 8080:80 --name welcome1 docker/welcome-to-docker
![imagewelcome](https://github.com/ivialvam/learning-docker/assets/97967090/d5fe324b-9a13-4282-b99e-917c2338d4a4)

2. Check whether the welcome1 container can run on the desired port

![port8080 80](https://github.com/ivialvam/learning-docker/assets/97967090/5d2d0c78-c113-4205-b375-6105a1432a12)

3. See logs of running welcome1 container in docker desktop
![welcome1](https://github.com/ivialvam/learning-docker/assets/97967090/7c3863f0-b99b-4869-9943-6faf2fa24d30)

4. Execute in exec command inside running welcome1 container cat usr/share/nginx/html/index.html

5. Stop container dan see lates logs
![welcome1 exit](https://github.com/ivialvam/learning-docker/assets/97967090/5f9e4a83-78b6-4953-a9da-290d1d1e48f9)



