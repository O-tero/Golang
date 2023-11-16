# Handling Authentication for our REST Services

# How simple authentication works

A client sends an authentication request to the server using user credentials.
The server takes those credentials and matches them with the credentials stored on the
server. If a match is successful, it writes something called a cookie in the response. This
cookie is a small piece of information that is transferred by the client for all subsequent
requests. Modern websites are being designed to be single-page applications (SPAs). In
those websites, static assets such as HTML and JavaScript files are served from a Content
Delivery Network (CDN) to render the web page initially. After that, the communication
between the web page and application server happens only through the REST API/web services.

A session records user communication in a given period of time whereby authentication information is stored in a cookie.
