# Goal 
To create a service when called on, responds back with the average grave and behavior.
After doing so will give a alert of either
- Red Average below 50
- Orange Average 50-70
- Green Average 70-100

## Steps

- Create repository logic that calls in all of grade and behavior rows; Returns a float64.
- Create repository logic that returns the number of students; Returns a float64
- Create service logic that calculates the average from the grave/behavior compared to the population; Which returns a string of the status.
- API just has one service that just sends back the status;But inside of this service it will need to call on 2 repo service.


## Solution
Just remembered I can just query postgreSQL and just get the average from there.
Don't need to over complicate this simple service.

