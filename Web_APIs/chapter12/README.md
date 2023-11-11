# Containerizing REST Services for Deployment

# ============================================

Containerization is required to avoid platform dependency during deployment of an application. To deploy an application
properly, we must prepare an ecosystem. That ecosystem consists of a web server, an application server, and a process monitor.

### Nginx Server

Nginx is a web server that can be a reverse proxy for a web application. It can also act as a load balancer when multiple instances of the server are running.
Supervisord makes sure that an application server is up and running in the event of a crash or a system restart.

Installation via Docker Container
`docker run --name nginxServer -d -p 80:80 nginx`

It starts a container with the name of nginxServer and serves it on port 80.
Accessing the containerized service
`http://localhost/` or `http://127.0.0.1/`

### Reverse Proxy Server

Server that holds the information regarding the original servers in it. It acts as the front-facing entity for the client request.
Whenever a client makes an HTTP request, it can directly go to the application server. However, if the application server is written in a programming language,
then you need a translator that can turn the application response into a response understandable by the clients. Common Gateway Interface (CGI) does the same thing.

#### Benefits of reverse proxy server(Nginx):

- It can act as a load balancer.
- It can provide access control and rate limiting.
- It can sit in front of a cluster of applications and redirect HTTP requests.
- It can serve a filesystem with a good performance.
- It streams media very well.

Nginx is also an upstream server i.e serves the requests from one server to the other.

### Important Nginx paths

| Type          |            Path             |                                                                                               Description |
| :------------ | :-------------------------: | --------------------------------------------------------------------------------------------------------: |
| Configuration |    /etc/nginx/nginx.conf    |                            This is the base Nginx configuration file. It can be used as the default file. |
| Configuration | /etc/nginx/sites-available/ |           If we have multiple sites running within Nginx, we can have a configuration file for each site. |
| Configuration |  /etc/nginx/sites-enabled/  |                                                         These are the sites currently activated on Nginx. |
| Log           |  /var/log/nginx/access.log  |                          This log file records the server activity, such as timestamps and API endpoints. |
| Log           |  /var/log/nginx/error.log   | This log file logs all proxy server-related errors, such as disk space, filesystem permissions, and more. |

### Using server blocks

Server blocks are the actual configuration pieces that tell the server what to serve and on which port to listen.
We can define multiple server blocks in the sites-available folder.
On Ubuntu, the location will be as follows:
`/etc/nginx/sites-available`

On Mac OS X, the location will be as follows:
`/usr/local/etc/nginx/sites-available`

Until we create a symlink from the sites-available to the sites-enabled directory, the configuration has no effect.
So, always create a symlink for sites-available to sites-enabled for every new configuration you create.

## Load balancing with Nginx

Load balancing is a process where the central server distributes the load to various servers based on certain criteria.

| Load-balancing   |                                                                   Method Description                                                                    |
| :--------------- | :-----------------------------------------------------------------------------------------------------------------------------------------------------: |
| Round Robin      |                         The incoming requests are uniformly distributed across servers based on the criteria of server weights.                         |
| Least Connection |                                 Requests are sent to the server that is currently serving the least number of clients.                                  |
| IP Hash          |  This is used to send the requests from a given client's IP to the given server. Only when that server is not available is it given to another server.  |
| Least Time       | A request from the client is sent to the machine with the lowest average latency (the time-to-serve client) and the least number of active connections. |

The first step in this process is to create an upstream cluster in the http section of the Nginx configuration file:
`http {
    upstream cluster {
        server site1.mysite.com weight=5;
        server site2.mysite.com weight=2;
        server backup.mysite.com backup;
    }
}`

Here, servers are the IP addresses or domain names of the servers running the same code.
We are defining an upstream called cluster here. It is a server group that we can refer to in our location directive.
Weights should be given in proportion to the resources available.
In the preceding code, site1 is given a higher weight because it may be a bigger instance(memory and CPU).
The backup server is given the backup parameter, which means that it will be used only when all the other servers are down.

### Rate limiting our REST API

We can limit the rate of access of our Nginx proxy server by rate limiting. This provides a directive called `limit_conn_zone`  
The format of it is this:
`limit_conn_zone client_type zone=zone_type:size;`

`client_type` can be one of two types:

- An IP address (limit requests from a given IP address)
- A server name (limit requests from a server)

`zone_type` also changes in correspondence to `client_type`. It takes values as per the
following table:

| Client type            | Zone type |
| :--------------------- | :-------- |
| $binary_remote_address | addr      |
| $server_name           | servers   |

Nginx has to save a few things in memory to remember the IP addresses and servers for rate limiting. 
The size parameter is the storage that we allocate for Nginx to perform its memory operations.
It takes values such as 8 m (8 MB) or 16 m (16 MB).