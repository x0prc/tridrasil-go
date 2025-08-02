## RateLimit Plugin

### Purpose:
Limits the number of actions (requests, commands, etc.) a user or client can perform in a given timeframe.

### Features
- Configurable rate limit policies (per user/API/key).
- Supports sliding window and token bucket algorithms.
- Returns warnings or errors on limit breach.

### Example
```
ratelimit:
  limit: 100
  window: 60  # seconds
```