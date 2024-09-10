# Sample Repository for Efficient Report Generation: Basic Approach
![3D89DB75-68F6-467B-BE14-573FCAD4C42D_4_5005_c](https://github.com/user-attachments/assets/f46f58cb-0519-41f4-89a9-c10953a1cd30)

This repository serves as a sample for a series of Medium blogs
1. [Efficiently Handling Both Fast and Slow Report Generation in Golang — Part 1: Separating Generation Task to a Consumer Service](https://medium.com/@chanokthorn6/efficiently-handling-both-fast-and-slow-report-generation-in-golang-part-1-the-problem-e83b1fa37f2b).
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

![image](https://github.com/user-attachments/assets/a2d3ddee-a531-4f05-8eab-9c23be695dfc)

## Installation

1. Clone the repository to your local machine.
2. Install the required dependencies by running `go mod tidy`.

## Usage

```
go run cmd/api/main.go
```

using APIs
```
curl --request POST 'http://localhost:3000?previous_days=1'
{
    "title": "Small Report",
    "content": "This is a small mock report."
}
```

This will start the application and generate sample reports based on the provided data.

## Contributing

Contributions to this sample repository are welcome. If you have any suggestions, improvements, or bug fixes, please submit a pull request.

## License

This sample repository is licensed under the [MIT License](LICENSE).
