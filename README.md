# Convert JSON string to Graphql 
> Convert JSON in string type (example below) to Grapql Input in Golang
## Setup
```
go get github.com/RizkyAnggita/converterGox
go mod tidy
```

## Code Examples
Example of usage:
```
jsonString := `
    {
      "_id": "615c05ddcb28473972a89134",
      "index": 0,
      "guid": "39237202-2fa0-4e43-ad6c-020e9fee7bba",
      "isActive": false,
      "balance": "$2,410.21",
      "picture": "http://placehold.it/32x32",
      "age": 31,
      "eyeColor": "green",
      "name": "Aguirre Bray",
      "gender": "male",
      "company": "NETPLAX"
    }`
res, err := goRizky.ConvertJSONGraphQL(jsonString)    

if err!=nil{
  return 
}

//Do something with res
fmt.Printf("Result --> %s", res)

    
```

## Output Example
```
Result --> {
      _id:  "615c05ddcb28473972a89134",
      index:  0,
      guid:  "39237202-2fa0-4e43-ad6c-020e9fee7bba",
      isActive:  false,
      balance:  "$2,410.21",
      picture:  "http://placehold.it/32x32",
      age:  31,
      eyeColor:  "green",
      name:  "Aguirre Bray",
      gender:  "male",
      company:  "NETPLAX"
    }
```
