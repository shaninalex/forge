
It should parse and "understand" how to work with this kind of expressions:

- response body.status = succeeded
- json body.order_id == {{ steps.checkout.response.body.order_id }}
- json response.status == 200
- json response.body.status == "PAID"
- json response.price > 0
- text body == "Access forbidden"
- header contain "AuthCookie="

> it's not designed properly yet. Just throwing ideas.

It also should know how to get values from previous steps to have dynamic tests.