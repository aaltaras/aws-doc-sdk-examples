/*
   Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
    "flag"
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

// CreateQueue creates an Amazon SQS queue
// Inputs:
//     sess is the
//     queueName is the name of the queue
// Output:
//     If success, the URL of the queue and nil
//     Otherwise, an empty string and an error from the call to CreateQueue
func CreateQueue(sess *session.Session, queueName string) (string, error) {
    // Create a SQS service client
    svc := sqs.New(sess)

    result, err := svc.CreateQueue(&sqs.CreateQueueInput{
        QueueName: aws.String(queueName),
        Attributes: map[string]*string{
            "DelaySeconds":           aws.String("60"),
            "MessageRetentionPeriod": aws.String("86400"),
        },
    })
    if err != nil {
        return "", err
    }

    return *result.QueueUrl, nil
}

func main() {
    queueNamePtr := flag.String("n", "", "The name of the queue")
    flag.Parse()

    if *queueNamePtr == "" {
        fmt.Println("You must supply a queue name (-n QUEUE-NAME")
        return
    }

    // Create a session that get credential values from ~/.aws/credentials
    // and the default region from ~/.aws/config
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    url, err := CreateQueue(sess, *queueNamePtr)
    if err != nil {
        fmt.Println("Got an error creating the queue:")
        fmt.Println(err)
        return
    }

    fmt.Println("URL for queue " + *queueNamePtr + ": " + url)
}