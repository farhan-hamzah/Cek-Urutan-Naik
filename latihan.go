package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var n, i, j int
	var b bool
	b = true
	fmt.Scan(&n)

	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}

	for i = 0; i < n; i++{
		for j = i+1; j < n; j++{
			if A[i] > A[j]{
				b = false
			}
		}
	}
	fmt.Print(b)
}