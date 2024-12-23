package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func buildGraph(input []string) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)

	for _, line := range input {
		nodes := strings.Split(line, "-")
		a, b := nodes[0], nodes[1]

		// Initialize maps if they don't exist
		if graph[a] == nil {
			graph[a] = make(map[string]bool)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]bool)
		}

		// Add bidirectional edges
		graph[a][b] = true
		graph[b][a] = true
	}

	return graph
}

func findTriangles(graph map[string]map[string]bool) int {
	triangles := make(map[string]bool)

	for a := range graph {
		if !strings.HasPrefix(a, "t") {
			continue
		}

		for b := range graph[a] {
			for x := range graph[b] {
				if graph[a][x] {
					nodes := []string{a, b, x}
					sort.Strings(nodes)
					key := strings.Join(nodes, ",")
					triangles[key] = true
				}
			}
		}
	}

	return len(triangles)
}

func bronKerbosch(graph map[string]map[string]bool, r, p, x map[string]bool, maxCliques *[][]string) {
	if len(p) == 0 && len(x) == 0 {
		// Found a maximal clique
		clique := make([]string, 0, len(r))
		for v := range r {
			clique = append(clique, v)
		}
		sort.Strings(clique)
		*maxCliques = append(*maxCliques, clique)
		return
	}

	// Choose pivot
	pivot := ""
	maxDegree := -1
	for v := range p {
		degree := 0
		for u := range p {
			if graph[v][u] {
				degree++
			}
		}
		if degree > maxDegree {
			maxDegree = degree
			pivot = v
		}
	}
	if pivot == "" {
		for v := range x {
			degree := 0
			for u := range p {
				if graph[v][u] {
					degree++
				}
			}
			if degree > maxDegree {
				maxDegree = degree
				pivot = v
			}
		}
	}

	// Process vertices
	candidates := make([]string, 0, len(p))
	for v := range p {
		if pivot == "" || !graph[pivot][v] {
			candidates = append(candidates, v)
		}
	}

	for _, v := range candidates {
		// New sets for recursive call
		rNew := copySet(r)
		rNew[v] = true

		pNew := make(map[string]bool)
		for u := range p {
			if graph[v][u] {
				pNew[u] = true
			}
		}

		xNew := make(map[string]bool)
		for u := range x {
			if graph[v][u] {
				xNew[u] = true
			}
		}

		bronKerbosch(graph, rNew, pNew, xNew, maxCliques)

		delete(p, v)
		x[v] = true
	}
}

// Helper function to copy a set
func copySet(s map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k, v := range s {
		result[k] = v
	}
	return result
}

// findMaximalCliques returns all maximal cliques in the graph
func findMaximalCliques(graph map[string]map[string]bool) [][]string {
	p := make(map[string]bool)
	for v := range graph {
		p[v] = true
	}

	var maxCliques [][]string
	bronKerbosch(graph, make(map[string]bool), p, make(map[string]bool), &maxCliques)

	// Sort cliques by size (largest first)
	sort.Slice(maxCliques, func(i, j int) bool {
		return len(maxCliques[i]) > len(maxCliques[j])
	})

	return maxCliques
}

func main() {
	var input []string
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	start := time.Now()
	graph := buildGraph(input)
	result := findTriangles(graph)
	elapsed := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", elapsed)

	start = time.Now()
	maxCliques := findMaximalCliques(graph)
	cliques := strings.Join(maxCliques[0], ",")
	elapsed = time.Since(start)
	fmt.Println("Part 2:", cliques, "Time:", elapsed)
}
