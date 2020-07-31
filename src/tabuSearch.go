package main

import (
    "fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

var tabu_list [][]int
var index_list = [][]int{{0,0},{0,1},{0,2},{0,3},{1,0},{1,1},{1,2},{1,3},{2,0},{2,1},{2,2},{2,3},{3,0},{3,1},{3,2},{3,3}}
var penalty int 
var penalty_old = 4 

func matPrint(X mat.Matrix) {
    fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
    fmt.Printf("%v\n", fa)
}

func choise_index(select_num int) []int {
	index := make([]int, select_num)
	for i := range index {
		index[i] = rand.Intn(4*4)
	}
	return index
}


func generate_near(NS *mat.Dense, S_good *mat.Dense) (*mat.Dense, int) {
	// 近傍解の作成
    // 近傍解を局所変換するので、要素を3つ選択する。（3というのは適当）
	  select_num := 3
	  index_num := choise_index(select_num)
	  chg_flag := true
	  // 変更する値
	  new_var := make([]float64, select_num)
	  for i := range new_var {
		  new_var[i] = rand.NormFloat64()
	  }

	  for i := range tabu_list {
		  if i == 7 {
			  break
		  }

		  for j := 0; j < select_num; j++ {
			  if index_num[j] == tabu_list[i][j] {
				  chg_flag = false
			  }
		  }
	  }

	  if chg_flag == true {
		  for i := range index_num {
			NS.Set(index_list[index_num[i]][0],index_list[index_num[i]][1], new_var[i])
		  }
		  penalty = 0
		  // i行の各要素の値を合計する。
		  //for i := range NS {
		  for i := 0; i < 4; i++ {
			  sum_col_ns := 0.0
			  for j := 0; j < 4; j++ {
				  sum_col_ns = sum_col_ns + NS.At(i, j)
			  }
			  // この範囲に各行列の合計値がなれば終了
			  if !(sum_col_ns < 1.1 && sum_col_ns >= 0.95) {
				  penalty = penalty + 1
			  }
		  }
		  if penalty < penalty_old {
			  S_good.Copy(NS)
		  } else {
			  tabu_list = append(tabu_list, index_num)
		  }
	  }
	  return S_good, penalty
}


func main() {
	matrix_num := 4
	// Generate a 6×6 matrix of random values.
	data := make([]float64, matrix_num*matrix_num)
	zero_mat := make([]float64, matrix_num*matrix_num)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	S := mat.NewDense(matrix_num, matrix_num, data)
	S_good := mat.NewDense(matrix_num, matrix_num, zero_mat)
	NS := mat.NewDense(matrix_num, matrix_num, zero_mat)

	NS.Copy(S)
	penalty := 0
	for i := 0; i < 100000000; i ++ {
		S_good, penalty = generate_near(NS, S_good)
		if penalty_old > penalty {
			penalty_old = penalty
			fmt.Printf("%v\n", penalty_old)
		}
		
		if penalty_old == 0 {
			matPrint(S_good)
			for i := 0; i < 4; i++ {
				sum_col_sgood := 0.0
				for j := 0; j < 4; j++ {
					sum_col_sgood = sum_col_sgood + S_good.At(i, j)
				}
				fmt.Printf("%v\n", sum_col_sgood)
			}
			break
		}
	}
}