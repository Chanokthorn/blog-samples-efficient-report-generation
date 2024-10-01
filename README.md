# Sample Repository for Efficient Report Generation - Part2: On-time Response with Redis Pub/Sub

<img src="https://github.com/user-attachments/assets/3edaceae-35f2-4494-8e70-c387c8ff5551" width="500">

This repository serves as a sample for a series of Medium blogs

1. [Handling Fast and Slow Report Generation in Golang — Part 1: Asynchronous System Using RabbitMQ](https://medium.com/@chanokthorn6/efficiently-handling-both-fast-and-slow-report-generation-in-golang-part-1-the-problem-e83b1fa37f2b)
2. [Handling Fast and Slow Report Generation in Golang - Part 2: On-time Response with Redis Pub/Sub](https://medium.com/@chanokthorn6/handling-fast-and-slow-report-generation-in-golang-part-2-on-time-response-with-redis-pub-sub-1e26bb53962e)


In this blog post, we will explore various techniques and best practices for generating reports efficiently in software applications. We will cover topics such as optimizing data retrieval, designing efficient report templates, and leveraging caching mechanisms to improve performance. 

Feel free to explore the code and accompanying documentation to gain insights into how to implement efficient report generation in your own projects.

## Table of Contents

- [Architecture](#architecture)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Architecture

<img src="https://github.com/user-attachments/assets/79616673-4c48-4a88-8e3b-241d1eedd068" width="600">

## Installation

1. Clone the repository to your local machine.
2. Install the required dependencies by running `go mod tidy`.

## Usage
Set up RabbitMQ and MongoDB with docker-compose
```
docker-compose up
```
Runing API Server
```
go run cmd/api/main.go
```
Running Report Generator Consumer
```
go run cmd/generator-consumer/main.go
```
using APIs
```
// user calls generate report
// curl --request POST 'http://localhost:3000?previous_days=5'
{
    "jobID": "ae097961-d601-426b-8b6d-fed18784e569"
}

// user calls get report shortly after generating, the report is not done yet
// curl --location 'http://localhost:3000/ae097961-d601-426b-8b6d-fed18784e569'
{
    "status": "processing"
}

// user calls get report again when the report is done
// curl --location 'http://localhost:3000/ae097961-d601-426b-8b6d-fed18784e569'
{
    "report": {
        "title": "Large Report",
        "content": "This is a large mock report with a lot of content..."
    }
}
```
### Demo
Fast report generation: returning report with response:
![short-demo](https://github.com/user-attachments/assets/5453e84e-d109-44d7-a34f-70a9b3fcf3d8)

Slow report generation: returning job ID in response for querying later:
![long-generate-demo](https://github.com/user-attachments/assets/b22a6fac-1228-4c07-9f16-951f28b4ed23)

![long-get-demo](https://github.com/user-attachments/assets/00c0b3c4-4082-4997-9688-6ec33b37c097)

This will start the application and generate sample reports based on the provided data.

## Contributing

Contributions to this sample repository are welcome. If you have any suggestions, improvements, or bug fixes, please submit a pull request.

## License

This sample repository is licensed under the [MIT License](LICENSE).
