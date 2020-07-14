# GOFuture

Implementation of CompletableFuture(java) in GO

## FutureGO.Future

### Struct 
* wg sync.WaitGroup

### Methods

#### RunAsync
Run the given function in goRoutines.
##### Parameters
* fn func(wg *sync.waitGroup)

#### Get
It blocks the call until result is finished

#### GetWithTimeout
It blocks the call for a specific time and returns
if call does not finish in given time

##### Parameters
* timeout: time.duration