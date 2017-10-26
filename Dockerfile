FROM golang:1.4.2-wheezy

# -----------------
# Install dependencies
# -----------------

RUN apt-get update && apt-get install -y build-essential

# -----------------
# Copy files over
# -----------------

RUN mkdir -p /mnt/server
ADD . /mnt/server
WORKDIR /mnt/server

# -----------------
# Install Go dependencies
# -----------------

RUN make dep-install

# -----------------
# Set the server
# -----------------

EXPOSE 4260
