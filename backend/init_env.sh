#!/bin/bash
echo "PORT = \":5000"\ > .env
echo "MONGO_URI = \"vaihda minut!\"" > .env
echo "MONGO_COLLECTION = \"beverages\"" >> .env

cd tests
echo "PORT = \":5000"\ > tests/.env
echo "MONGO_URI = \"vaihda minut!\"" > tests/.env
echo "MONGO_COLLECTION = \"test_beverages\"" >> tests/.env

cd ..
echo "Init done!, remember to change the MONGO_URI in .env and tests/.env files!"