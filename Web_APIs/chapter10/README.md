### GraphQL and Go

According to the official documentation, the definition of GraphQL is as follows:
`"GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, 
makes it easier to evolve APIs over time, and enables powerful developer tools."`

GraphQL provides a few features out of the box:
_Schema (a type system)_
_Versionless API_
_Schema to Code_

A GraphQL schema is a syntax for defining the boundaries of an API. The boundaries contain information about what server resources are exposed via the API.
Since it allows an on-the-fly update of the schema without failing all the clients, it helps us create a versionless API. GraphQL provides client and
server libraries for handling the resources defined in a schema.
GraphQL is a language, not a runtime. So, someone has to translate a GraphQL schema to the code that a programming language can understand.
A few GraphQL client and server libraries can do some automatic code generation from the schema definitions.

#### Over-fetching and under-fetching problems in the REST API

`Over-fetching` happens by an API when a server sends data that is not required by the client. The API is defined in advance and a client only needs to follow the API
documentation. The drawback of this is that bandwidth is wasted.

`Under-fetching` is where a response that's sent by an API server is not sufficient for making decisions for a client. That leads to the following conditions:
_A client has to make subsequent API calls to different endpoints to procure data._
_A client has to compute manually on top of procured data and then merge it._
This is a very inefficient approach as clients have very few resources compared to servers.
For example, if a mobile device has to compute costly operations because of under-fetching, the API has to change its strategy to provide sufficient data.
