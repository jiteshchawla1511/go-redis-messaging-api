# Go-redis-messaging-api
A Go lang microservice to interact with redis messaging queue and SFTP server

## Prerequisites

- [Go](https://golang.org/dl/) installed on your system.
- A running [Redis](https://redis.io/download) instance.
- A running [SFTP](https://hub.docker.com/r/atmoz/sftp) instance
- AWS SNS Service 

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/jiteshchawla1511/go-redis-messaging-api.git
   ```

2. Navigate to the project directory:
   
   ```bash
   cd go-redis-messaging-api
   ```

3. Install the necessary dependencies:
   
   ```bash
   go mod download
   ```

4. Run the Redis Server

   ```bash
   redis-server
   ```
5. Run the SFTP Docker file (using Docker App or terminal)
   ```bash
   docker run -p 22:22 -d atmoz/sftp foo:pass:::upload
   ```

5. Run the main file

  1) Terminal 1 
   ```bash
   cd go-redis-publisher
   go run main.go
   ```

  2) Terminal 2
  ```bash
  cd go-redis-sftp-sync
  go run main.go
  ```
  


## API Endpoints 

```bash
POST : "http://localhost:3000/submit"
```
## JSON DATA FORMAT 

```bash
  type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int64  `json:"age"`
	Location string `json:"location"`
}
```

Example 
```bash
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 30,
  "location": "New York"
}

```

## ENV FILE 

SFTP_USER=

SFTP_PASSWORD=

SFTP_HOST=localhost

REDIS_ADDR=localhost:6379

SMTP_FROM=<email>

SMTP_PASS=

SMTP_HOST=

SMTP_PORT=587

SNS_ARN=

AWS_REGION=

## END VERDICT

### You will recieve the data from your api to your email or phone number whichever you have configured in your aws sns service

Thankyou 😄
