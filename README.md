# Semgrep Demo

This repository includes a simple golang file, and a rules file for semgrep

Semgrep installation instructions can be found [here](https://semgrep.dev/docs/getting-started/#run-semgrep-locally)

## Running this example

```
semgrep --config /path/to/rules.yaml server.go
```

The output should result in something similar to the below
```
running 1 rules...
server.go
severity:warning rule:port-80-usage-detected: Usage of port 80 detected, web services should be served on port 443

18::80
ran 1 rules on 1 files: 1 findings
```

Changing the port to something besides 80 in the golang file, should cause the semgrep run to succeed

## Alternate Approaches

The semgrep rule library includes this [rule](https://semgrep.dev/editor?registry=go.lang.security.audit.net.use-tls), which guides users to use the TLS function call instead of the plaintext one.

The rule in the library makes sense to use over the demo, but I wanted to explore the ability to find the substring in the function call.

## More Information on Semgrep

I put together a brief overview of semgrep [here](https://dev.to/prince_of_pasta/supercharging-application-security-with-semgrep-3e1i)