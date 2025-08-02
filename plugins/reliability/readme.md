## Reliablity Plugin

### Purpose:
Enhances reliability for actions or API calls with retries, backoff, and failure monitoring.

### Features
- Automatic retries for failed attempts.
- Exponential backoff.
- Failure/event logging.

### Example
```
reliable:
  maxRetries: 3
  backoffStrategy: exponential
```