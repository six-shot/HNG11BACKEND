# Use the official Go base image
FROM golang:alpine 

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go .

# Build the Go executable
RUN go build -o ip-greeter

# Expose the port the app will run on
EXPOSE 8000

# Define the command to run when the container starts
CMD ["./ip-greeter"]
