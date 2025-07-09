# Tridrasil

**Three Modular plugins to enhance Yggdrasil-based networks with advanced features.**

## Features

- **Reputation & Blacklisting:** Track peer behavior, adjust reputation, and automatically blacklist abusive nodes.
- **Rate Limiting:** Token bucket algorithm to control request rates per peer.
- **Link Quality Metrics:** Collect latency and packet loss data, calculate routing cost for smarter path selection.
- **Reliable Delivery:** Sequencing, acknowledgements, retransmission, and fragmentation for robust message delivery.

## Usage

1. **Import the desired plugin(s) in your Go project:**
    ```
    import "tridrasil-plugins/plugins/reputation"
    import "tridrasil-plugins/plugins/ratelimit"
    import "tridrasil-plugins/plugins/linkquality"
    import "tridrasil-plugins/plugins/reliable"
    ```

2. **Integrate with your peer management, routing, or message handling logic.**
    - See each plugin’s `README.md` for API details and examples.

## Tests

- Each plugin includes unit tests.
- Run all tests from the project root:
    ```
    go test ./...
    ```
