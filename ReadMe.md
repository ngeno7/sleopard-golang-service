# SMSLeopard Engineering Challenge

## Database
 - The database credentials are contained on `.env.example` file
 - The database structure is contained on the `tables.sql` file.

## Queueing
Rabbit MQ is used for queueing campaigns.

Once the campaign is initiated, details are sent into a queue and can be picked by the background services for sending. On initiating, the status of the campaign can be set as `sending` till the queue is clear then can be set as `sent`.

The functionality of sending campaign is incomplete.

## Role of AI
I have used ChatGPT to optimize code debugging and how to implement concepts such as pointers, database configuration and RabbitMQ.



## How to Run
1. Install packages

```sh

go mod tidy

```
To install the packages

2. Import tables and seed data

Update the .env file accordingly based on .env.example

```sh

mysql -u username -p dbname < tables.sql

```

3. Run the Project
At the root of the project.

```sh
go run cmd/api/main.go
```
