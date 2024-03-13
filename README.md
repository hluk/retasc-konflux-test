Proof of concept for an Go+OpenShift project tested and deployed via [App
Studio](https://github.com/redhat-appstudio/docs.appstudio.io).

## Development

Enable pre-commit hook to lint and auto-format the code:

```
pre-commit install
```

## Build and Run

Test the app locally with [odo](https://odo.dev/docs/introduction) and
[minikube](https://minikube.sigs.k8s.io/docs/):

```
minikube start --driver=qemu2
odo dev
```

Or with odo and podman:

```
odo dev --platform podman
```

Or with go build:

```
go run .
```

## Testing

```bash
odo run test
# or
odo run --platform podman test
# or
go test ./...
# or
```

## TODO

- Run tests in pull requests and on merge to main branch
- Ensure majority of the code is covered by tests
- Some simple Jira API calls
- Healthchecks
