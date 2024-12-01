# lucide

A Go packaged version derived from [lucide icons](https://github.com/lucide-icons/lucide).

## Usage

``` go

import "github.com/BartAlcorn/lucide"

lucide.CirclePlus(lucide.Props{})

```

where Props are fairly self-explanatory

``` go
type Props struct {
  Size   string
  Color  string
  Fill   string
  Stroke string
  Class  string // extra TailwindCSS classes, e.g. text-blue-500
}
```
