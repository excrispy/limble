# limble

## Acceptance Criteria ##
At Limble, we want to track various costs for our customers, and one of those is labor cost. In our test problem, the labor cost is fairly simple: it’s the total of all the time worked by the workers, multiplied by their hourly wage. A worker can log time on a task several times, and many workers can work on the same task. This is captured by rows in the logged_time table.  
Changes Needed:  
JSON Objects for Pie Graphs  
We want to be able to make this information available to the frontend application for pie graphs via an HTTP server. The pie graphs can slice in two different ways:  
By worker: the total cost of that worker across all tasks and locations  
By location: the total labor cost for tasks tied to a given location  
Create some endpoints in the dummy server included in the repo that will give back these results in JSON format.  
Completion Status  
Tasks can be either complete or incomplete. When we do our API calls for the pie graphs, the client should be able to specify if they want to include completed tasks, incomplete tasks, or both.  
For example: if we’re doing average costs for tasks, incomplete tasks will throw off the average because they don’t give an idea of what the final cost will be.  
Filtering  
We should allow filtering on sets of worker IDs and/or location ID for both endpoints.  


## Notes
Run `docker compose up` to start the service and mysql db. Hard coded sample data will be generated.

Run `mysql -h 127.0.0.1 -P 3306 -u root -p` to log into the mysql cli.

Run `docker-compose down -v` to clean out the data.  

User name and password are root for testing purposes.

Sample cURL commands  
```curl --location 'http://localhost:8080/cost/by-worker?completed=0&worker_ids=1'```  
```curl --location 'http://localhost:8080/cost/by-location?completed=0&location_id=1'```

## Sample Tables

```
locations
+----+-------------+
| id | name        |
+----+-------------+
|  5 | Branch E    |
|  1 | Office A    |
|  4 | Site D      |
|  3 | Store C     |
|  2 | Warehouse B |
+----+-------------+

logged_time
+----+------------+---------+-----------+
| id | time_hours | task_id | worker_id |
+----+------------+---------+-----------+
|  1 |          2 |       2 |         1 |
|  2 |          3 |       3 |         1 |
|  3 |          2 |       4 |         1 |
|  4 |          4 |       5 |         1 |
|  5 |          3 |       1 |         2 |
|  6 |          1 |       3 |         2 |
|  7 |          4 |       4 |         2 |
|  8 |          2 |       5 |         2 |
|  9 |          2 |       1 |         3 |
| 10 |          2 |       2 |         3 |
| 11 |          3 |       4 |         3 |
| 12 |          3 |       5 |         3 |
| 13 |          1 |       1 |         4 |
| 14 |          2 |       2 |         4 |
| 15 |          3 |       3 |         4 |
| 16 |          2 |       5 |         4 |
| 17 |          3 |       1 |         5 |
| 18 |          3 |       2 |         5 |
| 19 |          2 |       3 |         5 |
| 20 |          4 |       4 |         5 |
+----+------------+---------+-----------+

tasks
+----+--------------------+-------------+-----------+
| id | description        | location_id | completed |
+----+--------------------+-------------+-----------+
|  1 | Clean windows      |           1 |         0 |
|  2 | Organize inventory |           2 |         1 |
|  3 | Restock shelves    |           3 |         0 |
|  4 | Repair roof        |           4 |         0 |
|  5 | Update computers   |           5 |         1 |
+----+--------------------+-------------+-----------+


workers
+----+---------------+-------------+
| id | username      | hourly_wage |
+----+---------------+-------------+
|  1 | john_doe      |       15.50 |
|  2 | jane_smith    |       16.75 |
|  3 | bob_johnson   |       14.25 |
|  4 | alice_brown   |       17.00 |
|  5 | charlie_davis |       15.00 |
+----+---------------+-------------+
```