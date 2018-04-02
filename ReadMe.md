# Commands

*Build*: sh build.sh

*Run the container and access the through localhost:8000*: docker run -p 8000:8080 -t viswanathct/simple-go

# ToDo

* Try passing the service name as an environment variable.

# Log

* 180321

  * 0616  Added a few more routes.

* 180322

  * 0029  Connected back to the service through it's own kubernetes-service-name as the identifier.

* 180402

  * 0000  The server now takes an optional command line argument, the port to listen.
  * 0101  Added the route grpc/greet to call a internal gRPC service (implemented elsewhere).
  * 1346  Bug Fixed: The service wasn't exposed (@ 0.0.0.0).
