Package _ordinal_ provides utility functions for converting numbers to their ordinal text form.

For example, when you want to provide 3 and get back "3rd".

Right now there is only one function, _ordinal.Itoa_, and using it looks like this:
```go
for i := -1; i < 1025; i++ {
  fmt.Println(ordinal.Itoa(i))
}

// this prints
//
// negative    // negative numbers always return this
// 0th         // today I learned 0th is a thing
// 1st
// 2nd
// 3rd
// 4th
// 5th
// 6th
// 7th
// 8th
// 9th
// 10th
// 11th        // accounts for teens
// 12th
// 13th
// 14th
// 15th
// 16th
// 17th
// 18th
// 19th
// 20th
// 21st
// ...
// 99th
// 100th
// 101st
// 102nd
// 103rd
// 104th
// 105th
// 106th
// 107th
// 108th
// 109th
// 110th
// 111th      // accounts for larger numbers ending in teens
// 112th
// 113th
// 114th
// 115th
// 116th
// 117th
// 118th
// 119th
// 120th
// 121st
// ...

```
