# Forge

Api testing tool.

``` 
$ forge --help
```

### Example:

```
$ go run app/main.go execute ./docs/example.yaml
Executing pipeline:  ./docs/example.yaml
[get_posts]: executing step
Process asserts...
[0] status eq 201: true
[get_posts]: done after 0.296527s
[create_user]: executing step
[create_user]: done after 0.165088s
[verify_user]: executing step
Process asserts...
[0] status eq 200: true
[1] body 0.name eq Leanne Graham: true
[2] body 0.id eq {{ create_user | id }}: false
```

### TODO:

- [ ] work with `text/plain`, `text/html` content types
- [x] process params in query
- [ ] process params in body
- [ ] process params in assertions
