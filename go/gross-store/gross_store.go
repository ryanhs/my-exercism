package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	v, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] += v
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Return false if the given unit is not in the units map.
	unitReduce, existsUnit := units[unit]
	if !existsUnit {
		return false
	}

	// Return false if the given item is not in the bill
	_, existsItem := bill[item]
	if !existsItem {
		return false
	}

	// Return false if the new quantity would be less than 0.
	if bill[item]-unitReduce < 0 {
		return false
	}

	// If the new quantity is 0, completely remove the item from the bill then return true
	// Otherwise, reduce the quantity of the item and return true.
	if bill[item]-unitReduce == 0 {
		delete(bill, item)
	} else {
		bill[item] -= unitReduce
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	// Return 0 and false if the item is not in the bill.
	qty, existsItem := bill[item]
	if !existsItem {
		return 0, false
	}

	return qty, true
}
