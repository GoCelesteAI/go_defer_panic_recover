# Go Defer, Panic & Recover

Source code for **Go Tutorial #20: Defer, Panic & Recover - Error Handling Patterns**

## Files

- `defer_basics.go` - Defer fundamentals and LIFO order
- `defer_cleanup.go` - Using defer for resource cleanup
- `panic_basics.go` - Understanding panic behavior
- `recover_basics.go` - Recovering from panics gracefully

## Running the Examples

```bash
go run defer_basics.go
go run defer_cleanup.go
go run panic_basics.go
go run recover_basics.go
```

## Key Concepts

### Defer
- Schedules a function call to run when the surrounding function returns
- Multiple defers run in LIFO (Last In, First Out) order
- Arguments are evaluated when defer is called, not when it runs

### Panic
- Stops normal execution flow
- Used for unrecoverable errors
- Prints stack trace and exits

### Recover
- Catches panics inside deferred functions
- Returns the panic value (or nil if no panic)
- Allows program to continue after a panic

## Watch the Tutorial

[Go Tutorial #20 on YouTube](https://youtube.com/@GoCelesteAI)

## License

MIT
