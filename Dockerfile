# Note: This image is also built in CircleCI, so limit references to internal repositories.
FROM golang:1.24.9

# Create app directory
WORKDIR /usr/src/app

# Copy app to directory
COPY . /usr/src/app

CMD ["make", "test"]
