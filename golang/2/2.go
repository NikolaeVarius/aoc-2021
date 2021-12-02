
package main
import (
    "fmt"
    "os"
    "log"
    "encoding/csv"
    "io"
    "strconv"
    // "sort"
    // "strings"
    // "time"
)

    
type position struct {
    currentX int
    currentY int
    aim int
}


func forward(pos position, delta int) position {

    newX := pos.currentX + delta
    newY := pos.currentY + (pos.aim * delta) 

    result := position{currentX: newX, currentY: newY, aim: pos.aim }


    return result
}

func increaseAim(pos position, delta int) position {

    result := position{currentX: pos.currentX, currentY: pos.currentY, aim: pos.aim + delta }

    return result
}

func decreaseAim(pos position, delta int) position {

    result := position{currentX: pos.currentX, currentY: pos.currentY, aim: pos.aim - delta}

    return result
}

func main() {
    var currentPosition position

    currentPosition.currentX = 0
    currentPosition.currentY = 0
    currentPosition.aim = 0


    f, err := os.Open("input.csv")

    if err != nil {
        log.Fatal(err)
    }

    r := csv.NewReader(f)

    for {

        record, err := r.Read()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatal(err)
        }

        direction := record[0]

        i, err := strconv.Atoi(record[1])
        value :=  i

        

        currentPosition = processDepth(direction, currentPosition, value)

        fmt.Printf("x")
        fmt.Printf("%+v\n", currentPosition.currentX)
        fmt.Printf("y")
        fmt.Printf("%+v\n", currentPosition.currentY)
        fmt.Printf("aim")
        fmt.Printf("%+v\n", currentPosition.aim)

        // fmt.Printf("%+v\n", direction)
        // fmt.Printf("%+v\n", value)
    }

}

func processDepth(direction string, subPosition position, distance int) (position) {
    var newPosition position

    if direction == "forward" {

        newPos := forward(subPosition, distance)
        newPosition = position{currentX: newPos.currentX, currentY: newPos.currentY, aim: newPos.aim }

    } else if direction == "down" {

        newPos := increaseAim(subPosition, distance)
        newPosition = position{currentX: newPos.currentX, currentY: newPos.currentY, aim: newPos.aim }


    } else if direction == "up" {

        newPos := decreaseAim(subPosition, distance)
        newPosition = position{currentX: newPos.currentX, currentY: newPos.currentY, aim: newPos.aim }
    }

    // fmt.Printf("%+v\n", newPosition.currentX)
    // fmt.Printf("%+v\n", newPosition.currentY)

    return newPosition
}