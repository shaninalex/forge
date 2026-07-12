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
[get_posts]: done after 0.195522s
[create_user]: executing step
[create_user]: done after 0.134993s
[verify_user]: executing step
Process asserts...
[0] status eq 200: true
[1] body 0.name eq Leanne Graham: true
[verify_user]: done after 0.029491s
Executing action completed after: 0.360139s
```

### TODO:

- work with `text/plain`, `text/html` content types
- process {{ ... }} in assertions and request body