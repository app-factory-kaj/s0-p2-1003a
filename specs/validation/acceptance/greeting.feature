Feature: Greeting

  @story-1
  Rule: A request naming a person receives a greeting addressed to them

    Scenario: Greeting a named person
      Given the greeter API is available
      When an API Consumer calls GET /hello with name "Alice"
      Then the response is a JSON greeting addressed to "Alice"

  @story-2
  Rule: A request with no name still receives a generic greeting rather than an error

    Scenario: Calling without a name
      Given the greeter API is available
      When an API Consumer calls GET /hello with no name
      Then the response is a generic JSON greeting

    Scenario: Calling with an empty name
      Given the greeter API is available
      When an API Consumer calls GET /hello with name ""
      Then the response is a generic JSON greeting
