
[OPENAPI - DOCS]
[README]
[github-ci, test, linter and build]
[ERROR HANDLER]
[IDEMPOTENCY]
[tilt-k8s]


Tests cases

Account-module
- Should return 404 if account is not found.
- Should return 422 if account ID is invalid.

Transaction-module
- If operation type is invalid should return 422
- If amout is null or zero should return 422
- Compra e saque should have a negative amount
- pagamento should be positive amount