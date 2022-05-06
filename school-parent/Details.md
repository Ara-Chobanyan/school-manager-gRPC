# The School Parent Service

This service will be called on by other micro-services if Delete, Edit, Add operation is required in the database.
It will take in the communication and query the contents of the communication inside of the database.

# Brain Storming
The parent is going just sit and listen to the other micro-services, and take in there request and in respond back by making changes to the database.
If everything goes well it should send back a ok status to the micro-service.
It shall use a hex architecture for allow easy injections into interfaces.

# The injection talk
So it seems like there is not need for http router which makes sense. For this is going be a gRPC server that's going listen for requests.
I need to create logic for the gRPC in a way that I can inject my repo dependency inside of it.


# Goals
- [x] Figure out to either use gRPC or REST. [Came to the conclusion going use gRPC for the entire project.]
- [x] Learn how to use gRPC to take in a request, and in response send a query with the request data.
- [x] Set up the models.
- [x] Work on the SQL interfaces and injection.
- [x] Make our SQL logic.
- [x] Figure out a structure that we inject our SQL logic in side of gRPC.
- [x] Get it to work.
- [ ] Create tests.
