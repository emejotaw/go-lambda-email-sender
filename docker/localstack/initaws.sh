zip -r main.zip main

aws --endpoint http://localhost:4566 secretsmanager create-secret \
--name aws-secret \
--secret-string '{"email": "youremail@email.com", "password": "yourpassword", "host": "smtp.gmail.com", "port": "587"}' 

aws --endpoint http://localhost:4566 iam create-role \
--role-name lambda-role \
--assume-role-policy-document "{"Version": "2012-10-17", "Statement": [
    {
        "Effect": "Allow",
        "Principal": {
            "Service": "lambda.amazonaws.com", 
            "Action": "sts:AssumeRole"
        }    
    }
]}"

aws --endpoint http://localhost:4566 sqs create-queue --queue-name aws-queue

aws --endpoint http://localhost:4566 lambda create-function \
--function-name aws-lambda \
--runtime go1.x \
--zip-file fileb://main.zip \
--timeout 10000 \
--environment Variables={"EMAIL_SECRET_ID=aws-secret,AWS_REGION=us-east-1,AWS_ENDPOINT=http://host.docker.internal:4566"} \
--handler main \
--role arn:aws:iam::000000000000:role/lambda-role

aws --endpoint http://localhost:4566 lambda create-event-source-mapping \
--function-name aws-lambda \
--batch-size 10 \
--event-source-arn arn:aws:sqs:us-east-1:000000000000:aws-queue