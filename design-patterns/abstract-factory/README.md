# Abstract Factory Design Pattern in Go

Abstract Factory is a creational design pattern, which solves the problem of creating entire product families without specifying their concrete classes.

Abstract Factory defines an interface for creating all distinct products but leaves the actual product creation to concrete factory classes. Each factory type corresponds to a certain product variety.

The client code calls the creation methods of a factory object instead of creating products directly with a constructor call (new operator). Since a factory corresponds to a single product variant, all its products will be compatible.

Client code works with factories and products only through their abstract interfaces. This lets the client code work with any product variants, created by the factory object. You just create a new concrete factory class and pass it to the client code.

If you can’t figure out the difference between various factory patterns and concepts, then read our [Factory Comparison](https://refactoring.guru/design-patterns/factory-comparison).


Conceptual Example

Say, you need to buy a sports kit, a set of two different products: a pair of shoes and a shirt. You would want to buy a full sports kit of the same brand to match all the items.

If we try to turn this into code, the abstract factory will help us create sets of products so that they would always match each other.

`iSportsFactory.go`: Abstract factory interface

```go
package main

import "fmt"

type ISportsFactory interface {
    makeShoe() IShoe
    makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
    if brand == "adidas" {
        return &Adidas{}, nil
    }

    if brand == "nike" {
        return &Nike{}, nil
    }

    return nil, fmt.Errorf("Wrong brand type passed")
}
```


`adidas.go`: Concrete factory

```go
package main

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
    return &AdidasShoe{
        Shoe: Shoe{
            logo: "adidas",
            size: 14,
        },
    }
}

func (a *Adidas) makeShirt() IShirt {
    return &AdidasShirt{
        Shirt: Shirt{
            logo: "adidas",
            size: 14,
        },
    }
}
```


`nike.go`: Concrete factory

```go
package main

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
    return &NikeShoe{
        Shoe: Shoe{
            logo: "nike",
            size: 14,
        },
    }
}

func (n *Nike) makeShirt() IShirt {
    return &NikeShirt{
        Shirt: Shirt{
            logo: "nike",
            size: 14,
        },
    }
}
```


`iShoe.go`: Abstract product

```go
package main

type IShoe interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type Shoe struct {
    logo string
    size int
}

func (s *Shoe) setLogo(logo string) {
    s.logo = logo
}

func (s *Shoe) getLogo() string {
    return s.logo
}

func (s *Shoe) setSize(size int) {
    s.size = size
}

func (s *Shoe) getSize() int {
    return s.size
}
```


`adidasShoe.go`: Concrete product

```go
package main

type AdidasShoe struct {
    Shoe
}
```

`nikeShoe.go`: Concrete product

```go
package main

type NikeShoe struct {
    Shoe
}
```

`iShirt.go`: Abstract product

```go
package main

type IShirt interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type Shirt struct {
    logo string
    size int
}

func (s *Shirt) setLogo(logo string) {
    s.logo = logo
}

func (s *Shirt) getLogo() string {
    return s.logo
}

func (s *Shirt) setSize(size int) {
    s.size = size
}

func (s *Shirt) getSize() int {
    return s.size
}
```


`adidasShirt.go`: Concrete product

```go
package main

type AdidasShirt struct {
    Shirt
}
```


`nikeShirt.go`: Concrete product

```go
package main

type NikeShirt struct {
    Shirt
}
```

`main.go`: Client code

```go
package main

import "fmt"

func main() {
    adidasFactory, _ := GetSportsFactory("adidas")
    nikeFactory, _ := GetSportsFactory("nike")

    nikeShoe := nikeFactory.makeShoe()
    nikeShirt := nikeFactory.makeShirt()

    adidasShoe := adidasFactory.makeShoe()
    adidasShirt := adidasFactory.makeShirt()

    printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

func printShirtDetails(s IShirt) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}
```

`output.txt`: Execution result

```txt
Logo: nike
Size: 14
Logo: nike
Size: 14
Logo: adidas
Size: 14
Logo: adidas
Size: 14
```
