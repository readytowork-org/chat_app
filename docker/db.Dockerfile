FROM mongo:latest

ENV MONGO_INITDB_ROOT_USERNAME=admin
ENV MONGO_INITDB_ROOT_PASSWORD=password123QWERTY

# COPY init-mongo.js /docker-entrypoint-initdb.d/
CMD ["mongod", ]




