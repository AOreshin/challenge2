package main

type pair struct {
  x, y int
}

func main() {
}

// visited should be already populated with occupied squares, so they appear as already visited
// starting speed is 0, 0
func shortestPath(visited map[pair]bool, graph map[pair][]pair, start, finish, speed pair, path []pair) []pair {
  if visited[start] {
    return path
  }
  path = append(path, start)
  if start == finish {
    return path
  }
  shortest := []pair{}
  for _, p := range adjacent(start, speed) {
    if !visited[p] {
      newPath := shortestPath(visited, graph, p, finish, speed, path)
      if len(newPath) > 0 {
        if len(shortest) == 0 || (len(newPath) < len(shortest)) {
          shortest = newPath
        }
      }
    }
  }
  return shortest
}

// adjacent contain all achievable squares with all possible speeds
func adjacent(point, speed pair) []pair {
  return nil
}

// calculating all possible speeds for a current speed
func speeds(p pair) []pair {
  speedLimit := 3
  result := []pair{p}
  if p.x + 1 <= speedLimit {
    result = append(result, pair{x: p.x + 1, y: p.y})
  }
  return result
}
