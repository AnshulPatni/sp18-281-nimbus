# MongoDB-Overview: 

MongoDB is a free and open-source cross-platform document-oriented database program. Classified as a NoSQL database program, MongoDB uses JSON-like documents with schemas. MongoDB is developed by MongoDB Inc., and is published under a combination of the GNU Affero General Public License and the Apache License. MongoDB is a cross-platform, document oriented NoSQL database that provides, high performance, high availability, and easy scalability. MongoDB uniquely allows users to mix and match multiple storage engines within a single deployment. This flexibility provides a more simple and reliable approach to meeting diverse application needs for data. Traditionally, multiple database technologies would need to be managed to meet these needs, with complex, custom integration code to move data between the technologies, and to ensure consistent, secure access.

MongoDB stores data as documents in a binary representation called BSON (Binary JSON). Documents that share a similar structure are typically organized as collections. You can think of collections as being analogous to a table in a relational database: documents are similar to rows, and fields are similar to columns.

## Advantages

* MongoDB stores data in flexible, JSON-like documents, meaning fields can vary from document to document and data structure can be changed over time.

* The document model maps to the objects in your application code, making data easy to work with

* Ad hoc queries, indexing, and real time aggregation provide powerful ways to access and analyze your data

* MongoDB is a distributed database at its core, so high availability, horizontal scaling, and geographic distribution are built in and easy to use

* Schema less − MongoDB is a document database in which one collection holds different documents. Number of fields, content and size of the document can differ from one document to another.


### Multimodel Architecture:

MongoDB uniquely allows users to mix and match multiple storage engines within a single deployment. This flexibility provides a more simple and reliable approach to meeting diverse application needs for data. Traditionally, multiple database technologies would need to be managed to meet these needs, with complex, custom integration code to move data between the technologies, and to ensure consistent, secure access. With MongoDB’s flexible storage architecture, the database automatically manages the movement of data between storage engine technologies using native replication.

* MongoDB’s flexible document data model presents a superset of other database models. It allows data to be represented as simple key-value pairs and flat, table-like structures, through to rich documents and objects with deeply nested arrays and sub-documents.

* With an expressive query language, documents can be queried in many ways – from simple lookups to creating sophisticated processing pipelines for data analytics and transformations, through to faceted search, JOINs and graph traversals.

* With a flexible storage architecture, application owners can deploy storage engines optimized for different workload and operational requirements.

### Replica sets
MongoDB maintains multiple copies of data called replica sets using native replication. A replica set is a fully self-healing shard that helps prevent database downtime. Replica failover is fully automated, eliminating the need for administrators to intervene manually.

The number of replicas in a MongoDB replica set is configurable: a larger number of replicas provide increased data availability and protection against database downtime (e.g., in case of multiple machine failures, rack failures, data center failures, or network partitions). Optionally, operations can be configured to write to multiple replicas before returning to the application, thereby providing functionality that is similar to synchronous replication.

### Load balancing
MongoDB scales horizontally using sharding. The user chooses a shard key, which determines how the data in a collection will be distributed. The data is split into ranges (based on the shard key) and distributed across multiple shards. (A shard is a master with one or more slaves.). Alternatively, the shard key can be hashed to map to a shard – enabling an even data distribution.

MongoDB can run over multiple servers, balancing the load or duplicating data to keep the system up and running in case of hardware failure.

## Concurrency Control in MongoDB

Concurrency control allows multiple applications to run concurrently without causing data inconsistency or conflicts. This can be achieved in our project with either of the two options: 

* Create a unique field that can have unique values during a write operation to track the changes. This prevents insertions or updates from creating duplicate data. A unique index is created on multiple fields to force uniqueness on that combination of field values.

* Specify the expected current value of a field in the query predicate for the write operations. The two-phase commit pattern provides a variation where the query predicate includes the application identifier as well as the expected state of the data in the write operation.


## Consistency in MongoDB

#### Monotonic Writes

MongoDB provides monotonic write guarantees for standalone mongod instances, replica sets, and sharded clusters. Suppose an application performs a sequence of operations that consists of a write operation W1 followed later in the sequence by a write operation W2. MongoDB guarantees that W1 operation precedes W2.

#### Real Time Order

For read and write operations on the primary, issuing read operations with "linearizable" read concern and write operations with "majority" write concern enables multiple threads to perform reads and writes on a single document as if a single thread performed these operations in real time; that is, the corresponding schedule for these reads and writes is considered linearizable.


# MongoDB on the AWS Cloud: 

MongoDB is an open source, NoSQL database that provides support for JSON-styled, document-oriented storage systems. It supports a flexible data model that enables us to store data of any structure, and provides a rich set of features, including full index support, sharding, and replication.

AWS enables us to set up the infrastructure to support MongoDB deployment in a flexible, scalable, and cost-effective manner on the AWS Cloud. This reference deployment will help us build a MongoDB cluster by automating configuration and deployment tasks. This Quick Start supports a self-service deployment of the MongoDB replica set cluster (version 3.2 or 3.4) on AWS.

If we want to set up a fully managed database service, we can use MongoDB Atlas instead of deploying this Quick Start. MongoDB Atlas creates a new VPC for your managed databases and automates potentially time-consuming administration tasks such as managing, monitoring, and backing up your MongoDB deployments.

## MongoDB Constructs
Here are some of the building blocks that are used in this reference deployment.

**Replica set**. Refers to a group of mongod instances that hold the same data. The purpose of replication is to ensure high availability, in case one of the servers goes down. This reference deployment supports one or three replica sets. In the case of three replica sets, the reference deployment launches three servers in three different Availability Zones (if the region supports it). In production clusters, we recommend using three replica sets (Primary, Secondary0, Secondary1).

All clients typically interact with the primary node for read and write operations. It is possible to choose a secondary node as a preference during read operations, but write operations always go to the primary node and get replicated asynchronously in the secondary nodes. If you choose a secondary node for read operations, watch out for stale data, because the secondary node may not be in sync with the primary node. For more information about how read operations are routed in a replica set, see the MongoDB documentation.

In a development environment, you can start with a single replica set and move to three replica sets during production.


# Up and Running MongoDB Docker:

Few commands that we can use to run MongoDB on containers:

#### MongoDB latest Docker pull: 
` docker pull mongo:latest `

#### Run MongoDB Docker: 
` docker run -d -p 27017:27017 --name mongodb mongodb `

#### Run Mongod w/ persistent/shared directory:
` docker run -d -p 27017:27017 -v <db-dir>:/data/db --name mongodb mongodb `

#### Run Mongo Bash:
` docker exec -it mongodb bash `



# Reference: 

https://www.mongodb.com/mongodb-architecture

https://docs.aws.amazon.com/quickstart/latest/mongodb/overview.html

https://docs.mongodb.com/manual/core/read-isolation-consistency-recency/

https://docs.mongodb.com/manual/tutorial/deploy-replica-set/