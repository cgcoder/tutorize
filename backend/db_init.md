# PostgreSQL Initialization Guide

## 1. Initialize the Data Directory

First, create a data directory and initialize it:

```bash
# Create a directory for your data
mkdir -p /path/to/data/directory

# Initialize the database cluster
/path/to/postgres/bin/initdb -D /path/to/data/directory -U postgres -W
```

## 2. Start PostgreSQL Server

```bash
# Start the server
/path/to/postgres/bin/pg_ctl -D /path/to/data/directory -l logfile start
```

## 3. Create a User and Database

Connect to the default 'postgres' database that was created during initialization:

```bash
/path/to/postgres/bin/psql -U postgres
```

Within the psql prompt, create your user and database:

```sql
-- Create a new user (role)
CREATE ROLE tutorize_user WITH LOGIN PASSWORD 'your_secure_password';

-- Make the user able to create databases (optional)
ALTER ROLE tutorize_user CREATEDB;

-- Create a new database owned by your user
CREATE DATABASE tutorize_db OWNER tutorize_user;

-- Grant privileges on the database to your user
GRANT ALL PRIVILEGES ON DATABASE tutorize_db TO tutorize_user;

-- Exit from psql
\q
```

## 4. Test Connection with New User

```bash
/path/to/postgres/bin/psql -U tutorize_user -d tutorize_db -h localhost
```

## 5. Configure for Remote Access (Optional)

Edit the following files in your data directory:

- `postgresql.conf` - Set `listen_addresses = '*'` to listen on all interfaces
- `pg_hba.conf` - Add entries to allow remote connections

## 6. Database Configuration for Your Application

For your Tutorize application, use these connection parameters:

```
Host: localhost
Port: 5432
Database: tutorize_db
User: tutorize_user
Password: your_secure_password
```

Use these details in your application's database connection configuration.
