# School-Manager
The purpose of this service is to be able to filter, throughout the database and find students that pass a certain condition.
So that the admin or whoever accesses this service can see students who are doing well or not.

### Student Behavior service
The service will find students with low grade and behavior level.
This service will also give the option to flag students as critical level,
to inform this students needs extra help or attention.


### Student Exceptional service
The service will find students with high grade and behavior.
This service will also give the option to flag students as exceptional.
Which could send out a email to this student if they wish to be a tutor.


## Plan 
- Create the Behavior and Exceptional service.
- Create a helper function for the Behavior and Exceptional services, that flags the students; To Avoid unreadable code.

## Brain Storm
- Have 3 repository functions

## Todo 

- [x] Create interfaces and models
- [x] Create repo logic for postgreSQL
- [ ] Create gRPC logic. 
- [ ] Need to figure out logic on Flag.
