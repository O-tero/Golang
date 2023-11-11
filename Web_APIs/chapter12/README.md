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

_It can act as a load balancer._
_It can provide access control and rate limiting._
_It can sit in front of a cluster of applications and redirect HTTP requests._
_It can serve a filesystem with a good performance._
_It streams media very well._

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
