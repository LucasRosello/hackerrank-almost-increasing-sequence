package main

func almostIncreasingSequence(sequence []int) bool {

	flag := false

	/*
		This program loops through the slice to
		find any element greater than the next
	*/
	for i := 0; i < (len(sequence) - 1); i++ {

		if sequence[i] >= sequence[i+1] {

			/*
				If it is the 2nd time it finds an element
				greater than the next, return false
			*/
			if flag {
				return false
			}

			flag = true

			/*

				If it is the 1st time it finds an element greater
				than the next, it checks if any of the two
				elements can be removed

				Explanation: If either of the 2 can be removed
				then the check can continue

				Ex: 1, 2, 9,3 ---> The 9 can be removed

				If none can be removed then it is impossible
				to get to an incremental sequence, return false

				Ex: 1, 9, 8, 3 ---> Neither 9 nor 8 can be removed

			*/
			if i > 0 && i+2 < len(sequence) {
				if sequence[i] <= sequence[i-1] || sequence[i] >= sequence[i+2] {
					if sequence[i+1] <= sequence[i-1] || sequence[i+1] >= sequence[i+2] {
						return false
					}
				}

			}
		}
	}

	return true
}
