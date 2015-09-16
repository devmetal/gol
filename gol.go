package gol

import(
  "os"
  "fmt"
  "bufio"
  "encoding/json"
)

const live rune = '*'
const dead rune = '.'

type LifeMatrix struct {
  n int
  matrix [][]bool
  after [][]bool
}

func (lm *LifeMatrix) JsonString() ([]byte, error) {
  return json.Marshal(lm.matrix)
}

func (lm *LifeMatrix) Print() {
  for i:=0; i<lm.n; i++ {
    for j:=0; j<lm.n; j++ {
      if (lm.matrix[i][j]) {
        fmt.Print(string(live))
      } else {
        fmt.Print(string(dead))
      }
    }
    fmt.Print("\n")
  }
  fmt.Print("\n")
}

func (lm *LifeMatrix) Iter() {
  matrixLen := len(lm.matrix)

  for i:=0; i<matrixLen; i++ {
    rowLen := len(lm.matrix[i])

    for j:=0; j<rowLen; j++ {
      if (i >= lm.n || j >= lm.n) {
        lm.after[i][j] = false
        continue
      }

      livings := lm.getLivingCells(i,j)
      if (lm.matrix[i][j]) {
        if (livings < 2) {
          lm.after[i][j] = false
        } else if (livings > 3) {
          lm.after[i][j] = false
        } else {
          lm.after[i][j] = true
        }
      } else if (livings == 3) {
        lm.after[i][j] = true
      } else {
        lm.after[i][j] = false
      }
    }
  }

  for i:=0; i<len(lm.matrix); i++ {
    for j:=0; j<len(lm.matrix[i]); j++ {
      lm.matrix[i][j] = lm.after[i][j]
    }
  }
}

func (lm *LifeMatrix) getLivingCells(i int, j int) int {
  vectors := [8][2]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}
  livings := 0
  for _, v := range vectors {
    x, y := v[0], v[1]
    if (i + x >= 0 && i + x < len(lm.matrix)) {
      if (j + y >= 0 && j + y < len(lm.matrix[i+x])) {
        if (lm.matrix[i+x][j+y]) {
          livings++
        }
      }
    }
  }
  return livings
}

func NewLifeMatrix(file *os.File, n int) *LifeMatrix {
  m := make([][]bool, 0)
  a := make([][]bool, 0)

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    mr := make([]bool, 0)
    ar := make([]bool, 0)

    for _, c := range line {
      if (c == live) {
        mr = append(mr, true)
        ar = append(ar, true)
      } else {
        mr = append(mr, false)
        ar = append(ar, false)
      }
    }

    m = append(m, mr)
    a = append(a, ar)
  }

  return &LifeMatrix{n, m, a}
}
