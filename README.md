# Digantara Scheduler

## Setup
1. Clone the **Digantara** repository  
2. At root (`Digantara`) run:
   1. `chmod +x wait-for.sh`
   2. `docker-compose up --build`

---

## Endpoints

### 1. **GET** `/jobs`
List all jobs from the database.

---

### 2. **POST** `/jobs`

**JSON Input**:
```json
{
   "cron": "*/10 * * * * *",          // required
   "type": "sms",                     // required
   "message": "cron task message",
   "name": ""
}
```
