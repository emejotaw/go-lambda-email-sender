# go-lambda-email-sender #

### A lambda which consumes from SQS and sends email ###

### Build the project ###
```sh
GOOS=linux go build main.go
```

### Execute lambda function ###
```sh
docker compose up -d
```
### Access the container ###
```sh
docker exec -it go-lambda-email-sender bash
```

```sh
aws --endpoint http://localhost:4566 sqs send-message --queue aws-queue --message-body '{"email": "jhondoe@gmail.com", "provider": "test", "type", "test"}'
```

### Check if message has stopped in queue ###
```sh
aws --endpoint http://localhost:4566 sqs get-queue-attributes --queue aws-queue --attribute-name All
```

### Invoke lambda directly ###
```sh
aws --endpoint http://localhost:4566 lambda invoke --function-name aws-lambda --payload '{"Records": [{"Body": "{\"email\": \"jhondoe@gmail.com\", \"provider\": \"test\", \"type\": \"test\"}"}]}' response.json --log-type Tail
```

### Get log group name ###
```sh
aws --endpoint http://localhost:4566 logs describe-log-groups
```


### Get log stream name ###
```sh
aws --endpoint http://localhost:4566 logs describe-log-streams --log-group /aws/lambda/aws-lambda
```

### Get log events ###
```sh
aws --endpoint http://localhost:4566 logs get-log-events --log-group-name /aws/lambda/aws-lambda --log-stream-name ${log_stream_name}
```
