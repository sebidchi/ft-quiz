## Requirements To Start app

Run make setup command to configure all needed stuff:
This is important as the setup step as it does the following:
- Spin up containers

```shell
$ make setup
```


To run tests just:

```shell
$ make test
```

Be sure that the hosts & ports added in your local files are reachable from your environment.

## Project structure

I'm following Hexagonal architecture and DDD. Notice that having the separation of the layers.
Having the domain layer in charge of the business logic and the application layer in charge of transporting the data.
Infrastructure layer is in charge of the persistence (Ports and Adapters).

I'm not using ValueObjects in most of the cases. The business logic is not to complex and by adding named types and VO's it'll be a overkill.

As you can see, I have a unit test for each use case. I normally do unit tests in that way, as all parts will be tested.
However, when something in domain has a lot of logic, I'm creating a test for each specific file, as you can see in UsersResults. 

Application is structured in contexts and modules. We have only one context and a couple of modules.

- Context: The app.
- Module
- - Quiz
- - UserAnswers

Quiz module is the one that will be used to allocating the quiz, having the use case of returning the questions and answers.
UserAnswers module is the one that will be used to store the answers and the use case of returning the results.

As you can see, I also created the Dependency Injection (DI) for the modules separatelly. Why? Because I'm not using the same DI for the whole app and I want to have a clearer separation of concerns.

Each module is responsible for its own services, routes, and use cases registration. (QueryBus/CommandBus)

I've created a `common.go` file to have the common services used by all modules.

## Entrypoint

The service the following entrypoint

- Api. `cmd/quiz/web/main.go` Is in charge of all sync http request 
- Client. `cmd/quiz/client/main.go` Is in charge of raising Cobra Client and communicating with the api



## Solution

Reading the challenge, I saw the following actors:

- Quiz

So that's why I've decided to create a module:

- UserAnswers. I consider that we should have separated the user answers from the quiz. This is because the answers are not related to the quiz (for this simple app).

The request/responses are in JSON. I wanted to use JSONAPI, but finally I decided to discard it.
I've added  middlewares to  controllers:

- Panic middleware. If a runtime panic, catched there.
- RequestId middleware.
- JsonSchema Validator

## Tests

We have two types of tests:

- Unit. Testing the use cases and mocking the infrastructure implementation dependencies. Using interfaces to achieve this. 
- Acceptance: Setting up the whole app. Services, asserting responses and more things.
- Let's discuss more about this in person.

## Client

I used Cobra as described in the challenge. 
I did in a very simple way, not validating responses from the api apart from errors. 
I used a selector to avoid users to be able to introduce wrong responses and having to have extra errors checking.

## Monitoring/Observability

- I normally will add Traces to log all the request and things that matters
- Logs. We have two kind of logs:
  - Warning. Things that happen in the app but are not a problem. We set monitors in DD, NR, ELK (wherever we have logs stored), having insightful ones per
  operation having thresholds and if the number of warnings is higher raising a alert. 
  - Critical. Critical means something is really causing problems to our application. This raise an alert inmediatly .
  - We are logging at the top of the application. We leave the error from the bottom to the infra layer and then we log errors while also returning http messages.
- With DD , NR, or Prometheus, I'd add things that are relevant for the behaviour of our app. like:
  - Response times
  - Messages accumulated in queues
  - Deadletters queues messages number
  - Variations in trends (ok response, 4XX responses, etcs)
  - Many more we can discuss in a conversation

## Documentation

Well, I had no time to add a proper swagger (openapi3), sorry. It's something I normally do, but lot of time invested already.

## Conclusion

This is a very simple app, but I think it's a good example of how to use the Hexagonal architecture and DDD. 
