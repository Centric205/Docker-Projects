# Deploying Services

In this project I will be deploying a service using **Docker Compose V3**.
Docker Compose is a tool used for defining and running multi-container Docker applications. 
It uses a YAML file to help us configure our application's services. 
After you've typed in your configurations, with a single command, you create and start all 
the services from your configuration.

To demostrate how it works, will be building a simple Python Wed application which uses the 
Flask framework and maintains a hit counter in Redis.

1. ## Add application and it's dependencies 

    ![Alt text](image.png)

2. ## Create a Dockerfile

    ![Alt text](image-1.png)

3. ## Create a Compose file and define your services

    ![Alt text](image-2.png)

4. ## Build and run your app with Compose

    ![Alt text](image-3.png)

5. ## Edit the Compose file to add a bind mount

    ![Alt text](image-4.png)

6. ## Re-build and run the app with Compose

    ![Alt text](image-5.png)

7. ## Application running

    ![Alt text](image-6.png)