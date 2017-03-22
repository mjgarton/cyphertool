# cyphertool
Simple utility for doing useful things with cypher statements

## Installation

Once you have [Installed Go](https://golang.org/doc/install), use:
```
go get github.com/utilitywarehouse/cyphertool
```

## Help

```
cyphertool --help
```

## Example usage
```
echo 'CREATE (ee:Person { name: "Emil", from: "Sweden", klout: 99 })' | cyphertool run
```

Multiple statements can be separated by `;`
