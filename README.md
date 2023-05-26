# 06_RMS-chi-db

BEST PRACTICES (done)
- REGISTER, LOGIN, LOGOUT
  -  authentication token for both register and login handlers for user experience
  -  password hashing
  -  logout no need for taking id from URL, logout should be from token 
- CONTEXT
  - can add context to middleware for user details that we can access in handlers
  - 
- ERROR HANDLING
  - create different error types
  - return specific error types from dbHelpers to handlers
  - generic function for error response to client from response writer
- QUERY & URL Params Differences
  - query params are optional
  - specific use case

-----------------------------------------------------------------------------------------------------------------------------------------

Remaining
- TRANSACTIONS
  - without wrapper
  - with wrapper 


