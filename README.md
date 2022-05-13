# School-Manager-gRPC

School-Manager is a group of microservices meant to be used to manage a school.
It is built with GO and gRPC using a hexagonal architecture (hopefully I implemented it correctly).
As well using PostgreSQL as the database.
Over all the code base is quite simple, the main goal of this project was, to tinker with gRPC and using a hexagonal architecture.

# The Services

#### School Data: 
The service pulls in data from the teachers,students and classrooms, as well giving the Population of students and teachers.

#### Global Grade Service
Finds the average grade and behavior of the school students, and sends back a signal of Red,Orange and Green. Based on how well the school is doing.

#### School Manager Service
This service flags students, who are in need so they can get more attention. As well flagging students who are doing really well, to be given rewards and be asked if they wish to 
tutor and help there class mates.

#### School Parent Service
The service adds,deletes and edits data on the database.
