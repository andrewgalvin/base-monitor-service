# Base Monitor Service

## Features

- Baseline repository for creating a multi-threaded application for crawling the internet

## Getting Started

### Prerequisites

- Go 1.20
- Database connection (repo uses MongoDB)

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/andrewgalvin/base-monitor-service.git
```

Navigate to the project directory

```bash
cd base-monitor-service
```

Install the necessary dependencies:

```
go mod tidy
```

## Running the service

To start the service, run:

```
go run cmd/base-monitor-service/main.go
```
