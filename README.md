# Price Compare V3

Price Compare V3 is a web application that allows users to compare prices of products across different stores. The application is built using Go, Echo framework, and Redis for data storage. The infrastructure is managed using Terraform and AWS.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Infrastructure](#infrastructure)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Prerequisites

- Go 1.16+
- Redis
- Terraform
- AWS CLI

### Clone the Repository

```bash
git clone https://github.com/evertonmj/price-compare-v3.git
cd price-compare-v3
```

### Install AWS CLI

To install the AWS CLI, follow these steps:

#### On macOS

```bash
brew install awscli
```

#### On Linux

```bash
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

#### On Windows

Download the installer from the [AWS CLI website](https://aws.amazon.com/cli/) and follow the installation instructions.

### Configure AWS CLI

To configure the AWS CLI, run the following commands:

```bash
aws configure set aws_access_key_id test
aws configure set aws_secret_access_key test
aws configure set region us-east-1
```

This will set up your AWS credentials and region for the project.

### Build the Application

```bash
cd app/src
go build -o price-compare-v3
```

### Run the Application

```bash
./price-compare-v3
```

## Usage

### Running the Server

To start the server, run the following command:

```bash
./price-compare-v3
```

The server will start on port 8080 by default.

## API Endpoints

### Health Check

- **GET** `/liveness` - Check if the service is alive
- **GET** `/readiness` - Check if the service is ready

### Products

- **POST** `/products` - Add a new product
- **GET** `/products/` - Get all products
- **GET** `/products/:id` - Get a product by ID
- **PUT** `/products/:id` - Update a product by ID
- **DELETE** `/products/:id` - Delete a product by ID

## Infrastructure

The infrastructure is managed using Terraform and AWS. The following resources are created:

- VPC
- Subnets (public and private)
- Internet Gateway
- Route Tables
- Security Groups
- EC2 Instances

### Deploying the Infrastructure
To deploy the infrastructure there are two options:

1. Install AWS CLI and run terraform commands from your machine.
2. Use CloudShell feature on AWS Console

#### Using AWS CLI

After installing AWS CLI, navigate to the `infra` directory and run the following commands:

```bash
cd infra
terraform init
terraform apply
```

#### Using AWS CloudShell

1. Open the [AWS Management Console](https://aws.amazon.com/console/).
2. Click on the CloudShell icon in the top navigation bar.
3. In the CloudShell terminal, clone the repository:

    ```bash
    git clone https://github.com/evertonmj/price-compare-v3.git
    cd price-compare-v3/infra
    ```

4. Initialize Terraform:

    ```bash
    terraform init
    ```

5. Apply the Terraform configuration:

    ```bash
    terraform apply
    ```

Follow the prompts to confirm the deployment.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.