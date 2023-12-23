package main

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	seq_map := map[string]bool{}

	for _, val := range seq {
		seq_map[val] = true
	}

	for val := range seq_map {
        println(val)
    }
}
