# Creating a simple Docker image that runs a Python script

# 1. Use an official Python runtime as a base/parant image
FROM python:3.8


# 2. Set working directory to /app
WORKDIR /app


# 3. Copy the current directory contents into the container
COPY . /app


# Run the Python script when the container launches
CMD ["python", "script.py"]
