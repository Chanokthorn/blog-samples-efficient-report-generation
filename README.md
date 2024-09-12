# Sample Repository for Efficient Report Generation: Separating Generation Task to a Consumer Service

![3D89DB75-68F6-467B-BE14-573FCAD4C42D_4_5005_c](https://github.com/user-attachments/assets/19d4aa1a-0b96-4c9c-ad5b-8a6cac547200)

This repository serves as a sample for a series of Medium blogs

1. [Efficiently Handling Both Fast and Slow Report Generation in Golang — Part 1: Separating Generation Task to a Consumer Service](https://medium.com/@chanokthorn6/efficiently-handling-both-fast-and-slow-report-generation-in-golang-part-1-the-problem-e83b1fa37f2b)
2. Report Generation Job Done Notification With Redis PubSub (stay tuned...)


In this blog post, we will explore various techniques and best practices for generating reports efficiently in software applications. We will cover topics such as optimizing data retrieval, designing efficient report templates, and leveraging caching mechanisms to improve performance. 

Feel free to explore the code and accompanying documentation to gain insights into how to implement efficient report generation in your own projects.

## Table of Contents

- [Architecture](#architecture)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Architecture

![image](https://github.com/user-attachments/assets/f9137506-ab1f-4f75-a62c-0f25ff2f0197)

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

This will start the application and generate sample reports based on the provided data.

## Contributing

Contributions to this sample repository are welcome. If you have any suggestions, improvements, or bug fixes, please submit a pull request.

## License

This sample repository is licensed under the [MIT License](LICENSE).
