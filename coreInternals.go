/*
# Core Internals of Go (Golang)

Go has its own concurrency runtime and doesn’t rely on OS-level threads like Node.js does with libuv.
It uses goroutines, channels, and a runtime scheduler to handle concurrency in a highly efficient way.

# Go Runtime Scheduler :
	Go uses a [work-stealing scheduler] to manage goroutines across multiple OS threads.

# Core Design: MPG Model

 `G` Goroutine : A lightweight unit of execution (user-level thread)
 `M` Machine : An actual OS thread
 `P` Processor : Manages and schedules goroutines to Machines

> Here a `P` runs multiple `G`s on an `M`, and the Go runtime maps these efficiently.

# Scheduler Works

1. `go myFunc()` → creates a new `G`.
2. Go runtime picks a free `P`, which is attached to an `M`.
3. That `P` executes the `G`.
4. If a `G` blocks (e.g., for I/O or sleep), the `M` is freed to pick another `G`.
5. If a `P` has no goroutines, it [steals] from other `P`s' queues.

> Highly efficient — can handle [millions of goroutines].


# Threads and System Calls

* Goroutines are [not 1:1 with threads] — Go uses [M\:N scheduling].
* A [small thread pool] is maintained by the runtime.
* If a syscall blocks, Go [parks] that goroutine and uses the thread for another one.

> Blocking sys calls don’t block the entire thread.


# Go vs Node.js Internals

 Feature               Go (Golang)                    Node.js (libuv)

Concurrency Model    -> Goroutines + Scheduler        | Event loop + Callbacks/Promises
Threading            -> M\:N Scheduling (G → Threads) | Single-threaded JS + libuv Thread Pool
Blocking I/O         -> Managed per goroutine         | Offloaded to thread pool
CPU-bound Tasks      -> Efficient with `GOMAXPROCS`   | Poor, unless using Worker Threads
System-level Threads -> Yes, managed by runtime       | Only via libuv (C++), not JS
Concurrency Support  -> Built-in (`go`, `chan`)       | Manual (`Promise`, `async/await`, etc.)

*/
