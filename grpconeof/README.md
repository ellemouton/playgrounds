# grpc `oneof`

This project demonstrates how a grpc client handles the case where it receives 
a `oneof` type that it does not yet know of. ie, the server is ahead of the client
in versions and knows of a new `oneof` option for a message (and possibly returns 
that new message to the client).

## See it for yourself!

1. Run the server:
      ```
   $ go run ./cmd/server
   ```
   
2. Run the client:

    ```
   // See how the client interprets a `oneof` message that
   // it already knows about.
   $ go run ./cmd/client old
   
   // See how the client interprets a `oneof` message that
   // it does not yet know about.
   $ go run ./cmd/client new
   ```
   
## Conclusion:

The client does _not_ error out in the case where it does 
not know about the new `oneof` option. Beautiful!