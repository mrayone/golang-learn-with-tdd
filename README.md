## Startup

Learning golang with quick book [quii](https://quii.gitbook.io/learn-go-with-tests/).

## Golang standard

- [x] Package managment.
- [x] Learned about integers.
- [x] Interactions.
- [x] struct, methods and interfaces.

## Golang test library

- [x] how create assertion helpers functions.
- [x] how write benchmarks.
- [x] how create tests example code (helpful to documentation).
- [x] how display diferente tests itentions with `tc.Run()`
- [x] how use [table driven tests](https://github.com/golang/go/wiki/TableDrivenTests)

## Maps

- [x] Create maps
- [x] Search items to maps
- [x] Update items in maps
- [x] Delete items from a map
- Learn more about
  - How to create errors that are constants
  - Writing error wrappers


## Mocking
Create mocks is really important, because we can test only business logic without
take care with third part systems. Our code would be more effective and take a fast
feedback.

> Without mocking important areas of your code will be untested. In our case we would not be able to test that our code paused between each print but there are countless other examples. Calling a service that can fail? Wanting to test your system in a particular state? It is very hard to test these scenarios without mocking.

> Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in slow feedback loops.

> By having to spin up a database or a webservice to test something you're likely to have fragile tests due to the unreliability of such services.

read more: [Test Dobule](https://martinfowler.com/bliki/TestDouble.html)

## Concurrency
Concurrency in go is great. We've to be able to create code that execute things in the same time. The syntax is nice and simplier.

### channels
Channel allow to communicate with caller function and write things in memory in the same time.
- Waiting groups configure that the function should be wait of the all goroutines finish
- Mutex tell to scheduler that the shared item of memory is open or lock to change, so the goroutine need to wait until the item is open to change.

### Select
Able that we receive data from multiple goroutines base on switch case statement. It's very helpful when we need to choice which data we want to use.
> Note: It's very common to create a case with a signal that abort channels to avoid a block.

### Httptest
A package that allow we interfaces a real net/http servers which is consistent.