<h1 align="center">
	Test Containers
</h1>

Example of using Test containers in Golang. Creating a **Nats** cluster with
containers and test nats cluster with Golang SDK while running unit tests.

## What are test-containers?

Test-containers is a package that makes it simple to create and clean up container-based 
dependencies for automated _integration/smoke_ tests.

We use _test-containers_ to run our dependent services like databases in a container
hence we won't need to mock them or use real external services. This makes our uint tests more
efficient and more reliable.

The clean, easy-to-use API enables developers
to programmatically define containers that should 
be run as part of a test and clean up those 
resources when the test is done.

## Example

In the following example, we are going to create a **Nats**
container and test it.
