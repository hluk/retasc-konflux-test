FROM registry.access.redhat.com/ubi9/go-toolset:1.20.12-3 as builder

# hadolint ignore=DL3045
COPY . .

RUN git config --global --add safe.directory "$PWD" \
    && go mod download \
    && go build -o ./main

FROM ubi9/ubi-micro:9.3

LABEL \
    name="retasc-konflux-test" \
    vendor="Red Hat developers" \
    summary="retasc-konflux-test proof oc concept" \
    description="Proof of concept" \
    maintainer="Red Hat, Inc." \
    license="MIT" \
    url="https://github.com/hluk/retasc-konflux-test" \
    io.k8s.description="retasc-konflux-test REST API" \
    io.k8s.display-name="retasc-konflux-test"

COPY --from=builder /opt/app-root/src/main /app/main

ENV PORT 8081
EXPOSE 8081

CMD ["/app/main"]
