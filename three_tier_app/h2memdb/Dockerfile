# Use an official Maven runtime as a parent image
FROM maven:3.8.5-openjdk-17 AS build

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the Maven project file
COPY ./ /usr/src/app

# Build the application
RUN mvn clean install

# Use an official OpenJDK runtime as a parent image
FROM openjdk:17.0-jdk-slim

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the application JAR file from the build stage
COPY --from=build /usr/src/app/target/h2memdb-0.0.1-SNAPSHOT.jar /usr/src/app

# Expose the port the app runs on
EXPOSE 8083

# Run the application
CMD ["java", "-jar", "h2memdb-0.0.1-SNAPSHOT.jar"]
