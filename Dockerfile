# Tardigrade runtime image.
#
# Build context is the repository root and it expects the OSS binary to
# already exist there (produced by `make buildoss`, which writes
# ./cockroachoss). Debian 11 still ships libncurses5, which the 2019
# binary links dynamically (via libedit); newer Debian releases
# dropped it.
#
#   make buildoss
#   docker build -t tardigrade:19.2 .
FROM debian:11-slim

RUN apt-get update && \
	apt-get install -y --no-install-recommends libc6 ca-certificates tzdata libncurses5 && \
	rm -rf /var/lib/apt/lists/*

RUN mkdir -p /cockroach
COPY cockroachoss /cockroach/cockroach
COPY build/deploy/cockroach.sh /cockroach/cockroach.sh

# Set working directory so that relative paths
# are resolved appropriately when passed as args.
WORKDIR /cockroach/

EXPOSE 26257 8080
ENTRYPOINT ["/cockroach/cockroach.sh"]
