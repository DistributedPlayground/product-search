# Product Search
The Products Search service is a [GraphQL](https://graphql.org/) API written in Go and containerized with [Docker](https://www.docker.com/). It's purpose is to facilitate customer queries for available products and product collections.
It is a fundamental part of our E-commerce platform within the [DistributedPlayground](https://github.com/DistributedPlayground) project. See the [project description](https://github.com/DistributedPlayground/project-description) for more details.

- [Service Architecture](#service-architecture)
- [Endpoint Description](#endpoint-description)
- [Running the Service](#running-the-service)

## Service Architecture
I chose a GraphQL API for this service because GraphQL allows flexible and easily extensible queries into complex data sets. This query flexibilty can lead to a reduction in overall network calls and volume of network data transferred relative to REST. 

## Endpoint Description

| Operation | Endpoint       | Description                                               | Parameters          | Return Type        |
|-----------|----------------|-----------------------------------------------------------|---------------------|--------------------|
| Query     | collections    | Retrieves a list of collections.                           | limit, offset       | [Collection!]!     |
| Query     | products       | Retrieves a list of products.                              | limit, offset       | [Product!]!        |
| Query     | collection     | Retrieves a specific collection by ID.                     | id                  | Collection!        |
| Query     | product        | Retrieves a specific product by ID.                        | id                  | Product!           |

### Request Parameters Examples:

- limit: Integer (e.g., 10)
- offset: Integer (e.g., 0)
- id: ID! (e.g., "abc123")


### Query Examples:
* Fetching a list of collections:
```graphql
query {
  collections(limit: 10, offset: 0) {
    id
    name
    description
  }
}
```

* Fetching a list of products:
```graphql
query {
  products(limit: 10, offset: 0) {
    id
    name
    description
    price
    quantity
    collection {
      name
    }
  }
}
```

* Fetching a specific collection:
```graphql
query {
  collection(id: "<collection-id>") {
    id
    name
    description
  }
}
```

* Fetching a specific product:
```graphql
query {
  product(id: "product-id-here") {
    id
    name
    description
    price
    quantity
    collection {
      name
    }
  }
}
```

## Running the Service
Follow the instructions in the [DistributedPlayground project description](https://github.com/DistributedPlayground/project-description).