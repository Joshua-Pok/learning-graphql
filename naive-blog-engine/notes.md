<!--markdownlint-disable-->

# N + 1 Query Problem

N + 1 is a problem typically associated with ORM and relationships. It refers to a situation where an application retrieves a collection of entities from a database along with asscoiated related entiteis. However, for each entity in the collection, an additional query is executed to fetch the related data, leading to N + 1 queries.


## How do we solve the N + 1 Query


### Eager Loading

### Batch Loading


# GraphQL Root Query

GraphQL Root Query is a special type in the GraphQL schema that lists all top level fields a client can query


We need to define it in our schema and it serves as a Menu for what clients can access



# Mutations

Mutations are the POST/PUT/DELETE of Graphql. Just like our schema has a RootQuery, we also have a RootMutation.


Mutations almost always require data, which we will extract from p.Args similar to queries



