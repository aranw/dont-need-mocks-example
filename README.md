Don't Need Mocks Example
========================

This repository offers a contrived example illustrating how mocks, when used with external systems such as Postgres, can introduce subtle bugs.

This example is a supplement to the blog post I wrote, which you can find [here](https://aran.dev/posts/you-probably-dont-need-to-mock/).

In our first example `withoutmocks` we are able to test our database logic and ensure it is working as we expect in both the database and service layer.

In our second example `withmocks`  we are using mocks to replace the database logic from our service layer which is unaware of a change that has been made at the database layer breaking the order in which the scores are returned.