# grpc errors

This playground contains a simple grpc server and client. The aim here was to
determine the best way of sending "matchable" errors across grpc since something
like `errors.Is()` would not work. 

## Running the example:
 1. Run the server:
   ```
   $ go run ./cmd/server
   ```
2. Run the client cli tool:
   ```
   $ go run ./cmd/client --help
   ```
   The client cli tool lets you choose between the following options so that you can observe how the errors from the
   server are interpreted for different methods of sending them.

   ```
   noerror  // The server sends no error back. Ie, a successful response.
   statusok  // The server sends an error with status code OK.           
   nonstatuserror // The server sends a non-status wrapped error.
   statuserror    // The server sends a status wrapped error.
   detailedstatuserror // The server sends a status wrapped error with the error type embedded in the status metadata.
   ```

## Conclusions:

Some conclusions I have drawn from this experiment: 
1. Always use `status.New(code, err.Err()).Err()` for returning errors to the
   client from the server side. This removes the need to convert the error to a
   grpc.Status on the client side if you want to extract the error Code.
2. The best way to match on errors across grpc is to create explicit error proto
   messages that can be matched on. These can then be added the message
`Status` using status.WithDetails()
