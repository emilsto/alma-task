#!/bin/bash
echo "PORT = \":5001\" > .env
echo "MONGO_URI = \"vaihda minut!\"" >> .env
echo "MONGO_COLLECTION = \"beverages\"" >> .env

echo "PORT = \":5001\" > tests/.env
echo "MONGO_URI = \"vaihda minut!\"" >> tests/.env
echo "MONGO_COLLECTION = \"test_beverages\"" >> tests/.env

echo "Init done!, remember to change the MONGO_URI in .env and tests/.env files!"
