# Go TCP Random Chat

## Overview

Go TCP Random Chat is an application built using Golang that allows users to chat with random individuals through TCP socket communication. This application does not use protocols like HTTP; instead, it relies solely on TCP socket communication to implement the chat functionality. Redis is used for data management, ensuring efficient and reliable performance.

The purpose of this project is to understand socket communication and to experience automating infrastructure provisioning using Terraform.

It is important to note that this application is not intended for production use.

## Requirements

- **Terraform**: Install Terraform from [terraform.io](https://www.terraform.io/downloads.html).

## Deploy your server using Terraform

0. **Prerequisites**:

   Before you begin, ensure that you have the following:

   - An AWS account with the necessary permissions to create resources.
   - An AWS access key ID and secret access key.
   - Terraform installed on your local machine.

1. **Clone the Repository**:

   Clone the repository to your local machine using the following command:

   ```bash
   git clone
   ```

2. **Navigate to the Terraform Directory**:

   Navigate to the `terraform` directory within the cloned repository:

   ```bash
   cd terraform
   ```

3. **Set Terraform variables**:

   Set the required Terraform variables in the `variables.tf` file. :

   When provisioning infrastructure using Terraform, you need to set the following variables:

   - **region**: The AWS region where your infrastructure will be provisioned (e.g., `us-west-2`).

   - **aws_access_key**: Your AWS access key ID for authentication.

   - **aws_secret_key**: Your AWS secret access key for authentication.

   - **app_name**: The name of the application, which in this case is "go-tcp-random-chat".

   - **app_port**: The port on which the application will run, default is 8888.

   - **app_image**: The Docker image for the application, default is "hjk1996/go-tcp-random-chat:latest".

   - **redis_node_type**: The instance type for the Redis node, default is "cache.t4g.micro".

   - **redis_num_nodes**: The number of Redis nodes, default is 1.

   - **min_capacity**: The minimum number of application instances to run, default is 1.

   - **max_capacity**: The maximum number of application instances to run, default is 3.

4. **Initialize Terraform**:

   Run the following command to initialize Terraform:

   ```bash
   terraform init
   ```

5. **Provision Infrastructure**:

   Run the following command to provision the infrastructure:

   ```bash
   terraform apply -auto-approve
   ```

   This command will create the necessary infrastructure components, including the VPC, subnets, security groups, Redis cluster, Network Load Balancer, and ECS cluster.

   You can find your server's endpoint in the Terraform output after the infrastructure has been provisioned.

## Commands

After you have connected to the server using telnet or your custom client, you can use the following commands

- **Create a New Room**:

  ```bash
  /new-room
  ```

  This command creates a new chat room.

- **Join a Room**:

  ```bash
  /join
  ```

  This command allows you to join an existing chat room randomly.

- **Leave the Current Room**:

  ```bash
  /leave
  ```

  This command lets you leave the current chat room.

- **Send a Message**:

  ```bash
  /msg <your message>
  ```

  Use this command to send a message to the other participant in the room.

- **Quit the Chat**:
  ```bash
  /quit
  ```
  This command disconnects you from the current chat room and ends the connection with the server.
