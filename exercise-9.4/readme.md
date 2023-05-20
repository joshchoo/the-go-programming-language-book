# Output

```sh
Creating pipeline...
Created pipeline of 1000000 goroutines.
Total stack in use: 2048491520 bytes
Stack use per goroutine: 2048 bytes
Sending 0
Done. Received value: 1000000. Took 365 ms
Send+Receives per ms: 5479
```

Results:
- Each Goroutine is allocated 2048 bytes of stack (note that this can grow and shrink as needed).