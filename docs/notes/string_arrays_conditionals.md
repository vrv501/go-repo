### Arrays & Slices
- Slices can be converted to arrays using [4]int(slice) -- this will copy first four elements from slice to new array. Since it's a copy, any change made to original slice will not affect the new array
- Arrays are passed by value
- However if you type-cats slice as (*[4]int)(slice) -- no it will be pointer to underlying array of original slice which means any changes to slice will affect this pointer variable as well
- Full Slice expression: slice[start:end:last-position-in-slice-to-share]. Last index will be last position in original slice that will be shared with this new slice. Capacity of new slice will be lastposition-start. So if new elements are appended to this new sice then it will no longer share underlying array of original slice. Instead it will be copied to new array pointed by slice. Which means original slice will be unaffaceted when capacity of new slice is increased

### Strings
- They are interpreted literals of slice of runes. Each rune in inturn represented by slice of bytes. So len(string) will give total number of bytes in string, not rune
- When you use index on a string, you are grabbing byte at that index not rune. 
- When you slice string, you're slicing slice of bytes which means if rune which is represented by more than one byte gets sliced then there are chances it will be incorrectly cut somewhere in between which might not be final expected result. Each index represents one byte
`[255 123 45 6 89]`
- You can typecast rune, byte as string(var). Similarly you can type-cast string as []rune(str) or []byte(str)
- for-range loop loops over rune in strings. So if you use `for i, v := range str`, the v is rune while i is starting index of rune. so dont be surprised if i is not consecutive as you loop over. You might need to type-cast string(v) to get string representation

### Maps
- Attempting to read from nil map returns zero value of valueType. Writing to nil map causes panic
- len(map) returns number of keys
- Not comparable
- Keys of map must be comparable(cant use slices, map)
- Attempting to delete from nil map or if key doesnt exist doesnt do anything
- Clearing a map using clear deletes all keys which means length of slice becomes zero not nil

### Structs
- No inheritance in go since no classes
- Struct instancess are comparable only when its attributes are comparable. If attributes arent comparable and you compare it will be runtime panic
- Different structs can't be compared even if they have same attributes, order in which they are defined and same types
- You can however type-cast a struct into another struct provided they have same attributes, order in which they are defined and type. Now since they are both same struct type you can compare them provided attributes are comparable


### Switch statements
- You can use break to break out of any case but it snot required as it happens implicitly
- You can use multiple values in case seperated by commas(case 3,4,5,6:). You can also have blank switch statement
- You can have expressions ot be evaluated in switch cases