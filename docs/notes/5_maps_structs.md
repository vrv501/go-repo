### Maps
- Attempting to read from nil map returns zero value of valueType. Writing to nil map causes panic
- len(map) returns number of keys
- Not comparable
- Keys of map must be comparable(cant use slices, map)
- Attempting to delete from nil map or if key doesnt exist doesnt do anything
- Clearing a map using clear deletes all keys which means length of slice becomes zero not nil
- However clearing a slice just sets elemnets until lenght to zero value

### Structs
- No inheritance in go since no classes
- Struct instancess are comparable only when its attributes are comparable. If attributes arent comparable and you compare it will be runtime panic
- Different structs can't be compared even if they have same attributes, order in which they are defined and same types
- You can however type-cast a struct into another struct provided they have same attributes, order in which they are defined and type. Now since they are both same struct type you can compare them provided attributes are comparable