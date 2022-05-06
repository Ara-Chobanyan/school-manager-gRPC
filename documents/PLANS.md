# School Manager 
The school manager will be able to manage.

# Services

## School Population Data Service [Bare bones done]

- Teacher service. 
  - Calculate there reviews and satisfaction.
- Student Service.
  - There grade and behavior in each class, and overall average.
- Classroom Service
  - Class rooms and the data inside of each class rooms.

## School Population Ratio Service !! [Going be merged into School Population Data Service; Also need to add some logic in the response]

- Ratio Service
  - Create a algorithm to find the ratio of teachers to students (example for every 1 teacher there are 30 students).
  - Flag overall school if they are low/high/medium on either students or teachers.

## School Population Management Service [Bare bones done]

- Student behavior service
  - Behavior problems and if students need to be suspended, expelled or given special treatment.
  - Create a sort of algorithm to sorts the entire school and finds the students with the lowest grade and behavior level and flag them.
- Student exceptional service
  - Finding students that are doing exceptionally well.
  - These students will be given rewards or enlist for tutelary program for students that are not doing well but have good behavior. 
  - Also manage to see if these students accept or refuse the tutelary program to flag them to avoid the system re-asking them.

## School Class Status Service !! [Going be merged with School Population Management]

- Flag teachers who seem to have bad students, to see if it's a student or a teacher issue.
  - Create algorithm to find out if the Teacher has low review and students with good behavior but low grades or vice versa.

## Global Grade Service

- Global grade of the school that calculate all the grade and behavior that finds the average that then create a status for the school.
  * RED: Global Average below 50
  * Orange: Global Average between 50-70 
  * Green: Global Average 70-100

## School Scheduler 

- Create a calender of the school day that include times of classes or events on certain days. That well shift the day in order to fit the events and classes.

## School Parent [done]

- Deletes,Adds,Edits students from the database.


## Tech Stack
- Go
  - Go rpc library  
  - Redis driver
  - PSQL driver = pgx
  - Google Calender API if I decide to use it
- PostgreSQL
  - For storing data
- Redis
  - Caching data to reduce Database queries
- Something AWS for hosting and for tinkering around with.

## Architecture for each service
- The goal is to create a hexagonal Architecture that way the software is separated into individual modules.
To create the ease of adding in new features, removing features or editing services with ease.

## Date Structures

|Student|Teacher|School|
|-------|-------|------|
|id INT|id INT|id INT|
|first_name STRING|first_name STRING|school_name STRING|
|last_name STRING|last_name STRING|teacher_pop INT|
|grade FLOAT|review FLOAT|student_pop INT|
|behavior FLOAT|satisfaction FLOAT|global_grade FLOAT|
|critical BOOL|critical BOOL|overall_satisfaction FLOAT|
|exceptional BOOL|Staff BOOL|critical BOOL|
|tutelary BOOL||pop_balance BOOL|
|satisfaction FLOAT|
|staff BOOL|
