# Asynchronous API Design

### Understanding sync/async API requests

A _synchronous_ request waits on the server to compute the result. On the other hand, an _asynchronous (async)_ request receives a response immediately with the information about the eventual result
A synchronous request is an HTTP request that blocks the server until the response is returned. The majority of the services on the web run in this fashion.
An asynchronous request returns with information that can be used to fetch the information of a process. This is due to the advent of distributed systems and loose coupling.
An asynchronous request instantly returns a response but not with the result. It issues a ticket for finding the status of the requested operation. A client can use that ticket (a different response) to check the status and the final result of the operation.
Example: The client is sends a request to the server and the server returns a response to the client. This response is not something the client can consume instantly. Long-running tasks/jobs can be made asynchronous by the server.
The client can then use the received response to find out the status of the job. Once the job is done, either the server can notify the client or the client can poll the result by looking at the status.

### Fan-in/fan-out of services

_fan-out_ refers to the process of forking a single request into multiple. Example: An e-commerce website integrating itself with a thirdparty payment gateway.
The website uses an API from the payment gateway to pop up the payment screen and enters security credentials. At the same time, the website may call another API called analytics to record the attempt of payment.
Another example is _MapReduce_. Map is a fan-in operation, while Reduce is a fan-out operation. A server can fan out a piece of information to the next set of services (API)
and ignore the result or can wait until all the responses from those servers are returned.

_Fan-in_ is an operation where two or more incoming requests converge into a single request. Example: An API aggregates results from multiple backend services and returns the result on the fly to a client.
A hotel price aggregator or flight ticket aggregator that fetches requested information about multiple hotels or flights from various data providers and displays them.

### Delaying API jobs with queuing

A server can have multiple workers running in parallel who can exhaust the contents of a queue and work on them. Whenever a client requests an operation through an asynchronous API,
the server can put that request in a job queue, and all the workers can pick up a task whenever their turn comes.
This approach can offload an API server and focus on its business logic instead of getting blocked on parallel/independent tasks such as sending emails, contacting third-party services, and so on.
A few use cases of queuing are as follows:

- Compress images and email the final result
- Automatic back pressuring (limiting the load on the server to predictable amounts)

### RabbitMQ, a powerful message queue

RabbitMQ implements a messaging protocol called Advanced Message Queueing Protocol(AMQP). It uses it to support worker queues.
It also supports many other data exchange patterns, such as the following:

- Publish/Subscribe
- Topic/Subscription
- Routing messages
- Remote Procedure Call (RPC)

Installing RabbitMQ on Docker `docker run -d — name dev-rabbit — hostname rabbitmq-dev -p 15672:15672 -p 5672:5672 rabbitmq:management`
RabbitMQ uses default port 5672 for its operations. You can change this using the initial settings for the Docker command.

### Communicating with RabbitMQ in Go

First, create a connection to dial to the broker. If the connection is successful, a Channel needs to be created out of the connection.
It has the API for performing operations on the message broker. Then, define a queue that messages are sent to. Finally, publish a message to the queue.
