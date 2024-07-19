package main


func SimulateAdderTuring() {

    xB := map[string][]string {
        "B,s1": {"(", "R", "s2"},
        "B,s2": {"1", "R", "s3"},
        "B,s3": {"1", "R", "s4"},
        "B,s4": {"+", "R", "s5"},
        "B,s5": {"1", "R", "s6"},
        "B,s6": {"1", "R", "s7"},
        "B,s7": {"1", "R", "s8"},
        "B,s8": {")", "R", "s9"},

        "B,s9": {"B", "L", "s9"},
        "),s9": {")", "L", "s9"},
        "1,s9": {"1", "L", "s9"},
        "+,s9": {"1", "R", "s10"},

       "1,s10": {"1", "R", "s10"},
       "),s10": {"B", "L", "s11"},

       "1,s11": {")", "R", "s12"},
        // iterate forward until stopped
       "B,s12": {"B", "R", "s12"},
    }
    tape := []string{"B", "B", "B", "B", "B", "B", "B", "B", "B", "B", "B", "B", "B", "B", "B"}
    simulate(tape, xB, 24)

}


