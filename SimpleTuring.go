package main

import "fmt"


func SimulateSimpleTuring() {

    xB := map[string][]string {
        "B,s1": {"X", "R", "s2"},
        "B,s2": {"B", "L", "s3"},
        "X,s3": {"B", "R", "s4"},
        "B,s4": {"B", "L", "s1"},
    }
    tape := []string{"B", "B"}
    simulate(tape, xB, 8)

}

func simulate(tape []string, instructions map[string][]string, cycles int) {
    head := 0;
    state := "s1"
    for i := 0; i < cycles; i++ {
        // visualisation of the turin machine
        track_head_icon := "    "
        fmt.Printf("%s: ", state)
        for j := 0; j < len(tape); j++ {
            fmt.Printf("%s", tape[j])
            if (j != head) {
                track_head_icon += " "
            } else {
                track_head_icon += "|"
            }
        }
        fmt.Printf("\n")
        fmt.Println(track_head_icon)

        // turing machine part
        instruction := instructions[tape[head] + "," + state]
        tape[head] = instruction[0]
        state = instruction[2]
        if (instruction[1] == "R") {
            head += 1
        } else {
            head -= 1
        }
    }

}


