package main

import "fmt"
import "strconv"

func main() {
	var S string
	fmt.Scan(&S)
	cnt := 0
	for i := 0; i<len(S)+1; i++ {
		for j := i+1; j<len(S)+1; j++ {
			n, _ := strconv.Atoi(S[i:j])
			if n % 2019 == 0 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

// def main():
//     S = input()
//     cnt = 0
//     for i in range(len(S)+1):
//         for j in range(i+1, len(S)+1):
//             if int(S[i:j]) % 2019 == 0:
//                 cnt += 1
//     print(cnt)

// if __name__ == '__main__':
//     main()
