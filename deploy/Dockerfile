# Tardigrade runtime image.
#
# Build context is the repository root and it expects the OSS binary to
# already exist there (produced by `make buildoss`, which writes
# ./cockroachoss). The slim base already carries libc6, ca-certificates
# and tzdata; libncurses5 (linked dynamically via libedit) is the only
# package this layer adds. Newer Debian releases dropped libncurses5,
# which is why this stays on 11.
#
#   make buildoss
#   docker build -f deploy/Dockerfile -t tardigrade:19.2 .
FROM debian:11-slim

RUN apt-get update && \
	apt-get install -y --no-install-recommends libncurses5 && \
	rm -rf /var/lib/apt/lists/*

RUN mkdir -p /cockroach
COPY cockroachoss /cockroach/cockroach
COPY build/deploy/cockroach.sh /cockroach/cockroach.sh

# Set working directory so that relative paths
# are resolved appropriately when passed as args.
WORKDIR /cockroach/

EXPOSE 26257 8080
ENTRYPOINT ["/cockroach/cockroach.sh"]
