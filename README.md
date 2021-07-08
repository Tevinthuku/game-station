## GameStation

A workspace that houses services that "mock" what PlayStation offers to its members.

### Why I built this.

I recently read & finished the book on clean architecture, and I love the idea of building services that let the business rules come first instead of the technologies, such as Databases, web frameworks, etc..

Ill focus on the business_rules/use cases of the product & then dive into how the services will be delivered to "users" on various platforms.

### Packages available

/pkg/gamestationnetwork

- This workspace/directory holds the user account's service. Playstation offers something similar, but in a broader sense. In that a Sony account which can be used for other Sony products can also be used in the Playstation network context. Since this is a small repo focused only on only Playstation like services, I chose to have gamestation_network house the user accounts.

/pkg/gamestationplus

- This service mocks the playstation plus service. In order to join service, you need to have a gamestation_network account & a subscription code needs to also be provided. This way you get to join `plus` and have an active subscription.

More to come

- [gamestationstore] - This will house a mock of the games provided on the Playstation store.
- Unit & integration tests - Clean architecture enables easy testing, so I'll work on having the various parts of the app tested.
- Integrate a web-framework - This will enable us to expose the app to users via REST.

https://github.com/golang/go/wiki/CodeReviewComments#go-code-review-comments
