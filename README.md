# Rest-GraphQL-gRPC-CRUD

This project is a CRUD application that uses three different approaches for building APIs: REST, GraphQL, and gRPC/gRPC-GW. It's orchestrated using Docker Compose for local setup. All 3 services utilises the same mongo database but a different approach to add/delete/update data in the db.

## Setup

    1. Clone this repository to your local machine.

    2. Run the following command to start the project:

        `docker-compose up --build`
        
### Access the frontends:

    REST Frontend: http://localhost:5000
    GraphQL Frontend: http://localhost:5001
    gRPC Frontend: http://localhost:5002

### Each frontend links to its respective backend:

    REST Backend: http://localhost:3001
    GraphQL Backend: http://localhost:8082
    gRPC Backend (via gRPC-GW): http://localhost:8081