# 06_RMS-chi-db

BEST PRACTICES (done)
- REGISTER, LOGIN, LOGOUT
  -  authentication token for both register and login handlers for user experience
  -  password hashing
  -  logout no need for taking id from URL, logout should be from token 
- CONTEXT
  - can add context of keys-values in request from middleware
  - ex. tokens, user details that we can access in handlers
- ERROR HANDLING
  - create different error types
  - return specific error types from dbHelpers to handlers
  - generic function for error response to client from response writer
- QUERY & URL Params Differences
  - query params are optional
  - specific use case
- SQL Query jmoiron
  - Get can also Returning columns while Inserting into database, alternative to Exec
- GLOBAL, CONSTANT and Other VARIABLES
  - use for DRY purpose and Readability
  - naming convention for readability  

-----------------------------------------------------------------------------------------------------------------------------------------

Remaining
- TRANSACTIONS
  - without wrapper
  - with wrapper 


