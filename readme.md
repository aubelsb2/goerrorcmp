# Go Error Comparer

Usage
```
func Test...(t *testing.T) {
    goerrorcmp.ErrorStringMatchesOrContains(t, err, each.ExpectedError)
}
```

Either can be nil.