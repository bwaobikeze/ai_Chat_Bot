# Project

Database as a service using [Neo4j](https://neo4j.com/docs/getting-started)

## Requirements

* [Docker](https://www.docker.com/products/docker-desktop/)
* [neo4j account](https://neo4j.com/cloud/platform/aura-graph-database/?ref=docs-nav-get-started)

## configuration
* navigate to [./server](./server)
* create a .env file
* add the following environment variables to the .env file: CONNECTION_URL = <add_your_neo4j_connection_url>, USERNAME = <add_your_neo4j_username>, PASSWORD = <add_your_neo4j_password>

## Start

To start this system, follow the steps below.

* navigate to [./scripts](./scripts)
* execute the [start_server.sh](./scripts/start_servers.sh) script
* run a HTTP post request to localhost:8081 with the following body: {"query":"<your_cypher_query>"}

## Components

This section describes the modular components of the system.

```
 ---------       
| client  |
 ---------       
   ^
   |
   v
 ---------            ---------
| server  | < --- > |  h2memdb  |
 ---------            ---------
   ^
   |
   v
 ----------
| database |
 ----------
```

* client - client application that has API for sending
  queries/requests to server

* server - web server serving requests

* h2memdb - in-memory cache layer

* database - actual graph database implementation

In the following subsection we describe each of the components in more
detail.

### Client

Client application is responsible for:
* accepting user input
* checking correctness of the query (lexing + parsing)
* sending the request to server (REST)
* accepting the response (json)
* caching the results
* visualizing the results

Client application will be written in several programming languages.

Input by the user will be a Cypher query for the database (which might
be valid or invalid). You can decide in what way to accept the input.
You can decide in what language to write this code (it can even be
Python).

Once you have the input, you should check if the input can be properly
parsed. You should use ANTLR to obtain lexer and parser; grammar file
is already available for Cypher
(https://s3.amazonaws.com/artifacts.opencypher.org/M23/Cypher.g4), so
you do not need to write your own. (We define N=4 in this case: 0-Java
lexer+parser, 1-C/C++ lexer+parser, 2-Go lexer+parser, 3-Python
lexer+parser.) If the input cannot be parsed, give a nice error (your
decision what the error should say) and ask for the next input. If the
input is correct, move to the next step.

Once the input is successfully parsed, you should prepare a request
for the server and send it (json). You can write request code in any
language you wish.

Upon receiving a response (json), you should store the response into a
local relational database. We define N=2 in this case: 0-h2 database,
1-sqlite database. The result should always be a single table.

Every response should be cached (you already stored it in a local db),
so that any future input/request that uses the same input do not need
to be send to the server. This can be written in any language. (Do not
forget that you might need to invalidate cache; feel free to be more
coarse-grained and invalidate everything if there is any modification
to the database in any query.)

### Server

Your server should be written in Go. The server will accept requests
and serve them. If the request has any error, appropriate message will
be sent to the client. If the request is valid, it will be sent to the
database, results will be accepted, packed, and sent to the user.

We leave to you to design communication protocol between the server
and database, e.g., log files, inter-process messages.

Client and server should communicate using REST. You have freedom to
define end points and arguments.

### H2memdb

A Java application with h2 in-memory database.

### Database

You should set up and use Neo4j as your actual database on the server.


## Testing

You should have tests for each part of your code; without tests, code
will be considered non-existent.


## Benchmarking

Graduate version only.

Write a bash script(s) that will collect benchmarking data for 100
queries. Each query should be run 100 times and averages should be
computed.

To start this system, follow the steps below.

* navigate to [./scripts](./scripts)
* execute the [benchmark.sh](./scripts/benchmark.sh) script

## Software

Your implementation should work under the following configuration:
* Linux (any recent distribution)
* N=2: 0 - Oracle Java 17 (https://www.oracle.com/java/technologies/downloads); 1 - Oracle Java 11
* N=2: 0 - Neo4j v5.x (cloud Graph Database Self-Managed community edition https://neo4j.com/deployment-center); 1 - Neo4j v4.x
* Go 1.18+
* h2 2.2.222
