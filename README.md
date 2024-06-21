# Go TCP Random Chat

## Overview

Go TCP Random Chat is an application built using Golang that allows users to chat with random individuals through TCP socket communication. This application does not use protocols like HTTP; instead, it relies solely on TCP socket communication to implement the chat functionality. Redis is used for data management, ensuring efficient and reliable performance. One of the key features of this project is the use of Terraform, which enables anyone to easily provision their own random chat server.

## Features

- **TCP Socket Communication**: The application exclusively uses TCP sockets for chat functionality, ensuring a lightweight and efficient communication protocol.
- **Random Chat Matching**: Connects users randomly to engage in conversations.
- **Redis Integration**: Utilizes Redis for data management, ensuring high performance and reliability.
- **Terraform Provisioning**: Allows users to provision their own random chat server easily using Terraform, making the deployment process straightforward and repeatable.

## Requirements

- **Golang**: Ensure you have Golang installed on your machine. You can download it from [golang.org](https://golang.org/dl/).
- **Redis**: You will need a Redis server instance. Installation instructions can be found at [redis.io](https://redis.io/download).
- **Terraform**: Install Terraform from [terraform.io](https://www.terraform.io/downloads.html).

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/go-tcp-random-chat.git
   cd go-tcp-random-chat
   ```

2. **Install Dependencies**:

   ```bash
   go mod download
   ```

3. **Configure Redis**:
   Ensure your Redis server is running and configured properly. Update the Redis connection details in the configuration file if necessary.

4. **Run the Application**:
   ```bash
   go run main.go
   ```

## Using Terraform to Provision the Server

1. **Navigate to the Terraform Directory**:

   ```bash
   cd terraform
   ```

2. **Initialize Terraform**:

   ```bash
   terraform init
   ```

3. **Apply the Terraform Configuration**:
   ```bash
   terraform apply
   ```
   Follow the prompts to provision your own random chat server.

## Usage

Once the application is running, users can connect to the chat server using a TCP client. Each user will be randomly paired with another user to start a conversation.

