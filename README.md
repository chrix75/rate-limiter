# Rate Limiter

This project implements a rate limiter in Go. It provides several implementations of rate limiting strategies,
including:

* **Counter Limiter:** This is a basic rate limiter that allows a certain number of calls within a given time period.
  It's suitable for scenarios where a burst of traffic is acceptable as long as the sustained rate remains within
  acceptable bounds.
* **Time Boxed Limiter:** Limits calls within a specific time window (e.g., per minute, hour, or day). Once the time
  window elapses, the limits reset. This is useful for controlling bursts of traffic while still allowing clients to
  proceed uninhibited after the initial burst subsides.
* **Day Limiter:** This limiter resets the limits on a daily basis. This strategy is useful for scenarios where there's
  a daily quota, such as free API calls, and that quota resets at the start of each new day.

## Design and Implementation

The rate limiter is designed with modularity and testability in mind. Key components include:

* **`RateLimiter` Interface:** Defines the common behavior for all rate limiter implementations, including
  `SetMaxCallsForClient` and `Allow`.
* **`LimitRepository` Interface:** Provides an abstraction for storing and retrieving client call counts. It includes
  implementations such as `InMemoryLimitRepository` and `ConcurrentLimitRepository`, supporting both testing and
  concurrent production environments.
* **`Timer` Interface:** Defines how to obtain the current time. Different timer implementations such as `ConstantTimer`
  and `DynamicTimer` provide flexibility for testing and real-world usage scenarios.

## Testing

Each rate limiter implementation is thoroughly tested using unit tests. The tests cover various scenarios, including:

* Allowing calls within limits.
* Refusing calls exceeding limits.
* Resetting limits based on time conditions.
* Correctness of limit updates.
