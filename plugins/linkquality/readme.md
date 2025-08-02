## Link Quality Plugin

### Purpose:
Evaluates and scores the quality/reliability of URLs and links based on predefined factors.

### Features
- Validates URLs (HTTP response, domain health).
- Checks for malicious content or broken links.
- Provides a link quality score.

### Example
```
linkquality:
  maxTimeout: 5  // seconds
  scoring:
    httpStatus: 30
    domainReputation: 70
```