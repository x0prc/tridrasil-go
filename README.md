# Tridrasil 🌿

**Three Modular plugins to enhance Yggdrasil-based networks with advanced features.**

[![Go Tests](https://github.com/x0prc/tridrasil-go/actions/workflows/go-tests.yml/badge.svg)](https://github.com/x0prc/tridrasil-go/actions/workflows/go-tests.yml)

## Features

- **Reputation & Blacklisting:** Track peer behavior, adjust reputation, and automatically blacklist abusive nodes.
- **Rate Limiting:** Token bucket algorithm to control request rates per peer.
- **Link Quality Metrics:** Collect latency and packet loss data, calculate routing cost for smarter path selection.
- **Reliable Delivery:** Sequencing, acknowledgements, retransmission, and fragmentation for robust message delivery.

## Usage
> [!NOTE]
> Modify the `plugin-name` according to the feature you are using / testing.
1. `cd plugins/plugin-name`
2. `go build -o plugin-name`

3. **Integrate with your peer management, routing, or message handling logic.**
    - See each plugin’s `README.md` for API details and examples.

## Tests
- Each plugin includes unit tests.
- Run all tests from the each plugin's directory:
```
cd tests
go test plugin-name
```
