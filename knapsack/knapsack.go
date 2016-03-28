package knapsack

import (
  "fmt"
  "math/big"
  "time"
)

type Problem struct {
  Items    []Item `json:"items"`
  Capacity int64  `json:"capacity"`
}

type Item struct {
  Name   string `json:"name"`
  Weight int64  `json:"weight"`
  Value  int64  `json:"value"`
}

type Solution struct {
  Instances    []int64
  Total_weight int64
  Total_value  int64
}

func SolveUnbounded(problem *Problem) Solution {
  defer timer(trace("SolveUnbounded"))
  return solveUnbounded(problem.Capacity, problem.Items)
}

func solveUnbounded(capacity int64, items []Item) Solution {
  if len(items) == 0 {
    return Solution{}
  }
  N := len(items)
  W := capacity
  weights := make([]int64, N)
  for i := range items {
    weights[i] = items[i].Weight
  }
  gcd_all := gcdAll(int64(W), weights)
  W /= gcd_all
  for i := range weights {
    weights[i] /= gcd_all
  }

  opt := make([]int64, W+1)
  used := make([]int, W+1)
  used[0] = -1
  for w := int64(1); w <= W; w++ {
    used[w] = used[w-1]
    best_value := opt[w-1]
    for n := 0; n < N; n++ {
      n_value := items[n].Value
      dw := w - weights[n]
      if dw >= 0 && (opt[dw] + n_value) > best_value {
        best_value = opt[dw] + n_value
        used[w] = n
      }
      opt[w] = best_value
    }
  }

  w := W
  instances := make([]int64, N)
  total_weight := int64(0)
  for ; w >= 0; {
    n := used[w]
    if n == -1 {
      break
    }
    instances[n] += 1
    w -= weights[n]
    total_weight += weights[n]
  }

  return Solution{instances, opt[W], total_weight}
}

func gcdAll(base int64, numbers []int64) int64 {
  gcd_all := base
  for i := range numbers {
    gcd_all = gcd(gcd_all, numbers[i])
  }
  return gcd_all
}

func gcd(x, y int64) int64 {
  return big.NewInt(0).GCD(nil, nil, big.NewInt(x), big.NewInt(y)).Int64()
}

func trace(s string) (string, time.Time) {
  return s, time.Now()
}

func timer(function string, start time.Time) {
  elapsed := time.Since(start)
  fmt.Printf("Function %v took %f seconds\n", function, elapsed.Seconds())
}
