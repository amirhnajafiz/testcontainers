<h1 align="center">
	Test Containers
</h1>

Example of using Test containers in Golang. Creating a **Nats** cluster with
containers and test nats cluster with Golang SDK while running unit tests.

We use _test-containers_ to run our dependent services like databases in a container
hence we won't need to mock them or use real external services. This makes our uint tests more
efficient and more reliable. Let's see an example.
