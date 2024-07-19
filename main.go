package main

import (
	"fmt"
)

func main() {

    xB := map[string][]string {
        "B,s1": {"X", "R", "s2"},
        "B,s2": {"B", "L", "s3"},
        "X,s3": {"B", "R", "s4"},
        "B,s4": {"B", "L", "s1"},
    }

    tape := []string{"B", "B"}
    head := 0;
    state := "s1"
    for i := 0; i < 8; i++ {
        // just to vizualise where the head is currently.
        track_head_icon := ""
        for j := 0; j < len(tape); j++ {
            fmt.Printf("%s",tape[j])
            if (j != head) {
                track_head_icon += " "
            } else {
                track_head_icon += "|"
            }
        }
        fmt.Printf("\n")
        fmt.Println(track_head_icon)

        result := xB[tape[head] + "," + state]
        tape[head] = result[0]
        state = result[2]
        if (result[1] == "R") {
            head += 1
        } else {
            head -= 1
        }
    }
}

