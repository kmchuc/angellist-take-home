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

type Response struct {
	Amounts map[string]float32
}

func ProrateCalculator(allocationAmount int, investors []Investor) (Response, error) {
	totalRequestedAmount := 0
	totalAverageAmount := 0
	//this is what we'll be returning
	proratedAmounts := make(map[string]float32)
	var maxAllocations []float32
	adjustedAllocationAmount := allocationAmount

	// calc sum of all investors requested amounts and sum of all investors average amounts
	for _, investor := range investors {
		totalRequestedAmount += investor.RequestedAmount
		totalAverageAmount += investor.AverageAmount
	}

	// adjust investor's amount they can invest using rate of their avg amount and sum of all investors
	for _, request := range investors {
		maxAllocations = append(maxAllocations, float32(allocationAmount)*(float32(request.AverageAmount)/float32(totalAverageAmount)))
	}

	// if investor requests to invest more than their adjusted amount
	// subtract the investors who are requesting less than their adjusted amount
	// from the totalAverageAmount and  adjustedAllocationAmount
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
			proratedAmounts[investor.Name] = float32(adjustedAllocationAmount) * (float32(investor.AverageAmount) / float32(totalAverageAmount))
		} else {
			proratedAmounts[investor.Name] = float32(investor.RequestedAmount)
		}
	}

	return Response{
		Amounts: proratedAmounts,
	}, nil
}
