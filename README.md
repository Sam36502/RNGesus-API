# RNGesus-API

One day, whilst at work, I was inspired to create the
holiest of random number generators which would allow
programmers to let God gently guide their systems.

## Usage
### Server / API

To host your own instance of the RNGesus API, you can either
build it, use the provided linux binary (sorry, windows users)
or run the docker container. (Public hosted example in progress)

The endpoints are as follows:

#### Endpoints

**`/v1/rand/float`:**
Generates and returns a random 64-bit floating point number

Params:
 - None

Response:

    {
        "num": 0.0000000000000000
    }

**`/v1/rand/int:`**
Generates and returns a random 64-bit integer

Params:
 - `min` (optional) Minimum allowed number
 - `max` (optional) Maximum allowed number

Response:

    {
        "num": 0
    }

**`/v1/pray/float:`**
Prays a certain number will show up more frequently

Params:
 - `num` (required) What float to pray for

Response:

    {
        "msg": "The monks are praying hard for '0.00000', Amen."
    }

#### Error Response
If an error occurs, it will return with status `666` and a message
object with some information about the error e.g.:

    {
        "msg": "Error: Invalid `num` parameter ('not-a-number') provided: strconv.ParseFloat: parsing \"not-a-number\": invalid syntax"
    }

### Client
I've also included a basic client into this repository for use in go.
Simply...

    go get github.com/Sam36502/RNGesus-API

... and then import "github.com/Sam36502/RNGesus-API/client"

Then you can create an `RNGesusClient`

```go
    package main

    import "github.com/Sam36502/RNGesus-API/client"

    func main() {
        rng := client.RNGesusClient{
            BaseUrl: "localhost:777",
        }
        fmt.Println("A random number granted by God himself:", rng.GetRandomFloat())
    }
```