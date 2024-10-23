# Mem-Deshed

Simple implementation of an in-memory database.

# RESP
Reference: 
- [Protocol Spec](https://redis.io/docs/latest/develop/reference/protocol-spec/)
- [RESP Implementation](https://www.build-redis-from-scratch.dev/en/resp-reader)

## Specs
Partial implementation for this project. 

> The `\r\n` is the protocol's terminator, which always separates its parts.
e.g. `some content 1\r\nsome content 2\r\n`

> The first byte in an RESP-serialized payload always identifies its type. Subsequent bytes constitute the type's contents. We categorize every RESP data type as either simple, bulk or aggregate. Simple types are similar to scalars in programming languages that represent plain literal values. Booleans and Integers are such examples.
e.g. `<identifier, simple | bulk | aggregate>`

| RESP data type | Category | First byte (ID)
| -------------------- | -------------- | -------------------
| Simple strings | simple | +
| Simple Errors | simple | -
| Integers | simple | :
| Bulk Strings | aggregate | $
| Nulls | simple | _
| Boolean | simple | #
| Double | simple | ,

**Simple Strings** Strings that do not contain `\n` or `\r` except for the termination sequence. [Reference](https://redis.io/docs/latest/develop/reference/protocol-spec/#simple-strings)

**Bulk Strings** General Strings. Requires a specified size.

## Examples
**Simple string response** `+OK\r\n`

**Invalid command** `-Invalid command\r\n`

**10000 and -10000 integer response** `:10000\r\n` and `:-10000\r\n`

**Multiline string** `$30\r\nhello world\nyou are amazing!\r\n`

**Nothing, null** `_\r\n`

**True/False** `#t\r\n` or  `#f\r\n`

**3.1415, -1/3, 1.2\*10^100** `,3.1415\r\n`, `,-0.333333333334\r\n`, and `,1.2E+100\r\n`
