<!--markdownlint-disable-->


# GraphQL Introduction

## What is GraphQL

GraphQL is a query language for our APIs. It also serves as a runtime for fufilling queries. It is typically transport agnostic but typically served with HTTP



A GraphQL Query only asks for the data it needs. The Query is nested and when executed, can traverse down related objects. This allows us to make a single request for two types of data.

Whenever a query is executed against a graphql server. It is validated against a type system. Every GraphQL service defines types in a **GraphQL Schema**. 


Examplee of a type Schema

```Go 

type Person {
    id: ID!
    name: String
    birthYear: String
    eyeColor: String
    gender: String
    hairColor: String
    height: Int
    mass: Float
    skinColor: String
    homeworld: Planet
    species: Species
    filmConnection: PersonFilmsConnection
    starshipConnection: PersonStarshipConnection
    vehicleConnection: PersonVehiclesConnection
    created: String
    edited: String
}



```


## GraphQL Design Principles

1) Hierchical 

Fields are nested within other fields and the query is shaped like the data it returns


2) PRoduct Centric

GraphQL is driven by the data needs of the client and the language and runtime that ssupport the client


3) Strong Typing

A GraphQL server is backed by a GraphQL type system


4) Introspective

GraphQL language is able to query the server's type system


## Benefits Of GrapQL over REST


When we make requests with REST APIs we are limited by the API sepcification. The API will return all the data from that endpoint, regardless of whether we need it or not. This causes **Overfetching**


The inverse is also true where we may sometimes need to make multiple requests to get all the information we need. This is called **underfetching**


## Creating GraphQL objects

GraphQL objects are type specifications

We use graphql.newObject to define a new type. Each field needs a Type: graphQl.String and optionally a resolve function


Resolve function is a function that is executed when the field is queried


A field can have a Args property which we access via res.Args


Because res.Args returns an interface{} we need to use Go's type assertion 

eg:
```Go


id := p.Args["id"].(string)
```


## Designing a Schema

In GraphQL, instead of looking at APIs as a collection of REST endpoints, we look at it as a collection of types.



### Connections and Lists

In GraphQL we can define fields that return lists of any other graph ql type

[string] defines a list of strings

| List declaration | definition | example |
| --------------- | --------------- | --------------- |
| [int] | a list of nullable | Item3.1 |
| [int!] | non nullable  | Item3.2 |
| non nullable list of nullable integer | [int]! | Item3.3 |
| Item1.4 | Item2.4 | Item3.4 |


Basically ! means not nullable


## Lists and Slices

**graphql.NewList**: This is a type wrapper. We take an existing type like "bookType" and pass it into this function. This tells graphql to expect an array of those objects rather than a single one


Slice Resolution:

Whenever a resolver returns a go slice, the engine automatically iterates over each element applying bookType definition to every item in the slice


# Field Level Resolvers


In Production, data is usually normalized. We do not store entire author inside every book, we can move the link logic out of the data structures and into the graphql execution layer


Field Level Resolvers allow us to resolve a specific field by looking up based on parent object's data.



# Context and Mock Authentication


Context is the bucket of data that travels along with the query to every single resolver in the tree

When we call graphql.Do we can pass a context.Context object. 

We can then access this context inside res.Context


We can ensure that context includes a secret token before returning any fields
