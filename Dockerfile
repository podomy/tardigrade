# Tardigrade runtime image.
#
# Build context is the repository root and it expects the OSS binary to
# already exist there (produced by `make buildoss`, which writes
# ./cockroachoss). The image only adds a current Debian userland:
# libc6 (dynamically linked by the cgo binary), ca-certificates and
# tzdata. glibc is backward compatible, so a binary built on the
# period toolchain runs fine here.
#
#   make buildoss
#   docker build -t tardigrade:19.2 .
FROM debian:12-slim

RUN apt-get update && \
	apt-get install -y --no-install-recommends libc6 ca-certificates tzdata && \
	rm -rf /var/lib/apt/lists/*

RUN mkdir -p /cockroach
COPY cockroachoss /cockroach/cockroach
COPY build/deploy/cockroach.sh /cockroach/cockroach.sh

# Set working directory so that relative paths
# are resolved appropriately when passed as args.
WORKDIR /cockroach/

EXPOSE 26257 8080
ENTRYPOINT ["/cockroach/cockroach.sh"]
