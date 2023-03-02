package util

type Allocate struct {
	AllocationAmount int        `json:"allocation_amount"`
	InvestorAmounts  []Investor `json:"investor_amounts"`
}

type Investor struct {
	Name            string `json:"name"`
	RequestedAmount int    `json:"requested_amount"`
	AverageAmount   int    `json:"average_amount"`
}

func AllocationCalculator(allocationAmount int, investors []Investor) (map[string]float32, int, error) {
	totalRequestedAmount := 0
	totalAverageAmount := 0
	// this is what we'll be returning
	proratedAmounts := make(map[string]float32)
	var maxAllocations []float32
	adjustedAllocationAmount := allocationAmount

	// sum of all investors requested amounts and sum of all investors average amounts
	for _, investor := range investors {
		totalRequestedAmount += investor.RequestedAmount
		totalAverageAmount += investor.AverageAmount
	}

	// adjust investor's amount they can invest using rate of their avg investment
	for _, request := range investors {
		maxAllocations = append(maxAllocations, float32(allocationAmount)*(float32(request.AverageAmount)/float32(totalAverageAmount)))
	}

	// if investor requests to invest more than their adjusted amount
	// subtract the investors who are requesting less than their adjusted amount
	// from the totalAverageAmount and adjustedAllocationAmount
	for index, investor := range investors {
		if float32(investor.RequestedAmount) <= maxAllocations[index] {
			totalAverageAmount -= investor.AverageAmount
			adjustedAllocationAmount -= investor.RequestedAmount
		}
	}

	for index, investor := range investors {
		// if investor's request amount > than their adjusted avg,
		// prorate using adjusted total avg amount and adjusted avg
		if totalRequestedAmount > allocationAmount && float32(investor.RequestedAmount) > maxAllocations[index] {
			proratedAmounts[investor.Name] = ProrateCalculator(adjustedAllocationAmount, investor.AverageAmount, totalAverageAmount)
		} else {
			// if total investor's requests <= allocationAmount and
			// investor's request is less than adjusted avg
			// just add the investor's original requested amount
			proratedAmounts[investor.Name] = float32(investor.RequestedAmount)
		}
	}

	return proratedAmounts, 200, nil
}

func ProrateCalculator(allocationAmount int, avgAmount int, totalAvg int) (proratedAmount float32) {
	return float32(allocationAmount) * (float32(avgAmount) / float32(totalAvg))
}
