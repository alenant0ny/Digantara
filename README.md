Setup:
1. Clone the Digantara repository
2. At root(Diganatara) run:
   1. chmod +x wait-for.sh
   2. docker-compose up --build


Endpoints:
1. GET /jobs
   List all jobs from DB

2. POST /jobs

   JSON INPUT
      {
         "cron": "*/10 * * * * *",          //required
         "type": "sms",                     //required
         "message": "cron task message",
         "name": ""
      }
   As of now the scheduler only accepts cron expression from the input. This needs to be updated to accept high  level information like "everyday", "monday", "2pm", etc and translate it to cron expression in the program logic.
   Creates a record in the DB for newly created jobs.

6. GET /jobs/id/:id
   List job by id. There are two ids. Job ID assigned by cron and Auto incremented primary key ID by db. Here the Primary Key ID is used instead of Job ID issued by cron. Job ID is saved in the DB.

Startup:
On application startup, the program executes any existing jobs saved in the db, so that it acts like restarting jobs that were supposed to run if the app had shut down unexpectedly. 
