# ugodict
Urban Dictionary wrapper in GO

##Example
```
//Initialize client
client:= ugodict.GetClient()

//Get definitions
result, err, err2 := client.DefineByTerm("Eugene")

//Check null
if result != nil {
    
    //Print data
    definition := result[0]
    fmt.Println("ID: " + strconv.Itoa(definition.DefId))
    fmt.Println("Author: " + definition.Author)
    fmt.Println("Word: " + definition.Word)
    fmt.Println("Definition: " + definition.Definition)
    fmt.Println("Example: " + definition.Example)
    fmt.Println("Thumbs up: " + strconv.Itoa(definition.ThumbsUp))
    fmt.Println("Thumbs down: " + strconv.Itoa(definition.ThumbsDown))
    fmt.Println("Permalink: " + definition.PermaLink)
    fmt.Println("Written on: " + definition.WrittenOn)
}
if err != nil {
	fmt.Println(err)
}
if err2 != nil {
	fmt.Println(err2)
}
```
