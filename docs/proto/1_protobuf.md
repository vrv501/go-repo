- RPC's require XDR format to store data sent/received from different systems. Protobuf's are one of popular XDR frameworks that help define how to do this
- Serialisation format with language & machine specific generated code used with structured typed data
- Can be used for ephemeral network transfer or long-term storage

## Compatability promise
- Old code can read new data with below constructs:
  - Ignore any newly added fields
  - Set default values for deleted fields
  - deleted repeated fields will be empty
- New code can read old data. Those fields which are absent will be set to reasonable default values

## Scalar DataTypes
- double, float, `[u]int[32|64]`,`sint[32|64]` ,`[s]fixed[32|64]`, bool, string, bytes
- uint[32|64]: if always positive
- int[32|64]: if both +ve and -ve, but mostly +ve
- sint[32|64]: if mostly -ve(can also be used for +ve)
- fixed[32|64]: if large +ve numbers(>2^28, 2^56)
- sfixed[32|64]: if large -ve numbers(<-2^28, <-2^56)
- string, bytes: maxlen is 2^32
- deserailised enums in go can store unknown value, so check if its valid first in go
- int32, uint32, int64, uint64, and bool are all compatible both ways with each other
- sint32 and sint64 are compatible both ways with each other
- string and bytes are compatible both ways with each other(as long as they are valid UTF-8)
- fixed32 is compatible with sfixed32
- fixed64 with sfixed64


## Ground Rules
- Do not rename fields once in production
- Reserve field numbers. Even if field is deleted, do not use that number for some other field
- Use plural names for repeated fields
- Generally avoid changing field types
- When creating rpc, support partialUpdates, not full replaces. Client shud send only those fields they want to update and server should only change those. This also means more rpc to update a specific field 
- Use strings for id 
- dont encode any strings in application
- Prefer repeated msgs over scalars, enums

## How to delete fields
- First delete refrences of field you are going to delete from client code
- Then delete the field in prootbuf and reserve field-number & field-name


## Field presense(Proto3)
- For string, numeric, enum, bytes if optional is set then explicit presense can be tracked
- For singlular messages it is always tracked
- For maps(optional is N/A), repeated(optional is N/A) not tracked
- For field in oneof's(optional is N/A) it is tracked 


## Other types
### Any
- "google/protobuf/any.proto" - google.protobuf.Any

### Oneof
- Oneof can be checked to see which field has been set using `case()` or `WhichOneOf()`
- You cannot have map or repeated fields in oneOf. However a message containing repeated can b e field of OneOf
- oneOf cannot be repeated
- Dont add/remove fields in oneOf due to backward compatability issues 

### maps
- key can be integral, string(but not double, float, bytes)
- not allowed enum, proto msgs
- values can be anything other than maps
- maps cannot be repeated


## Options
- [Options](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto)
- No useful options for enum types, enum values, oneof fields, service types, and service methods
- `int32 old_field = 6 [deprecated = true];` to say a **field** is deprecated

## Well-known types
- https://protobuf.dev/programming-guides/dos-donts/#well-known-common
- https://protobuf.dev/reference/protobuf/google.protobuf/ 
- googleapis/googleapis/type

## Go specific libraries
- grpc/go grpc, status, codes, [metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md), gooelais/rpc/errdetails
- grpc-go-middleware
- [keepalive](https://pkg.go.dev/google.golang.org/grpc/keepalive?utm_source=godoc)
- Implementing four kinds of grpc using go: https://grpc.io/docs/languages/go/basics/


## Misc
- grpc reads and writes are thread safe however when streaming, do not attempt to concurrent read or concurrent write, but concurrent read & write is ok(not multiple reads or writes but read & write can be done concurrently)
- status.Error returns error. You can add additional details to the error using status.WithDetails. Clients can convert error to status.Status and then get details using sttaus.Details
- add metadata to context `AppendToOutgoingContext`
- SendMetadata must be called before sending request or response and shud only be done once
- https://github.com/super-linter/super-linter 
- https://github.com/nilslice/protolock
- https://github.com/grpc/grpc-go/blob/master/examples/features

## HTTP-to-gRPC codes
```
200 - 0
400(bad req) - 3(invalid_argument),9(failed_precondition),11(out_of_range)
401(unauthorized lol) - 16(unauthenticated)
403(forbidden) - 7(permission_denied)
404 - 5(not_found)
409 - 6(already_exists), 10(aborted)?
429 - 8(resource_exhausted)
499 - 1(cancelled) - req cancelled by client
500 - 15(data_loss? - why shud we inform client what???), 2(unknown), 13(internal - this is better)
503 - 14(unavailable)
504 - 4(deadline_exceeded - set by client)
```

Below will never returned by library:
INVALID_ARGUMENT
NOT_FOUND
ALREADY_EXISTS
FAILED_PRECONDITION
ABORTED
OUT_OF_RANGE
DATA_LOSS





