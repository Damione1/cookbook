#!/bin/bash

# Check if the database container is already running
db_container=$(docker-compose ps -q db)

if [ -z "$db_container" ]; then
  # Start the database container
  echo "Starting the database..."
  docker-compose up -d db

  # Set a flag to indicate that the script started the container
  started_by_script=true
else
  echo "Database container is already running."
  started_by_script=false
fi

# Wait for the database to be ready
echo "Waiting for the database to be ready..."
while ! docker-compose exec -T db pg_isready; do
  echo "Database not ready yet, waiting..."
  sleep 2
done

# Generate the models
echo "Generating models..."
docker-compose run --rm web sqlboiler --debug psql

if [ "$started_by_script" = true ]; then
  # Stop the database container
  echo "Stopping the database..."
  docker-compose down
else
  echo "Database container was already running, not stopping it."
fi
