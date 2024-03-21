package question4

func FindMinGroceries(landDevelopment string) int {
	totalA := 0
	totalE := 0
	res := -100
	communityUpdate := []byte(landDevelopment)

	for i := range communityUpdate {
		if communityUpdate[i] == 'A' {
			totalA++
		}
		if communityUpdate[i] == 'E' {
			totalE++
		}
	}

	if totalA > totalE {
		return res
	}

	count := 0

	for i := 1; i < len(communityUpdate)-1; i++ {
		if communityUpdate[i] == 'E' {
			count++
			communityUpdate[i] = 'G'
		}
	}

	if count != 0 {
		res = count
	}
	return res
}
