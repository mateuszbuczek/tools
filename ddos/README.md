Put file 'request.json' with following structure inside folder as executable:
```
{
  "number_of_calls": "100",
  "method": "GET",
  "url": "https://some.url.com",
  "body": {
    "username": "randomString(4,10)",
    "password": "string"
    ...
  }
}
```

max number_of_calls = 4,294,967,296

Available functions:

- randomString(from,to) - generate string of size (from, to)
 with charset "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

