package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	eachLayerTime := time

	if time == 0 {
		eachLayerTime = 2
	}

	return len(layers) * eachLayerTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	// For each noodle layer in your lasagna, you will need 50 grams of noodles.
	// For each sauce layer in your lasagna, you will need 0.2 liters of sauce.
	gr := 0
	liter := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			gr += 50
		}
		if layers[i] == "sauce" {
			liter += 0.2
		}
	}

	return gr, liter
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	multiplier := float64(portions) / 2

	for i := 0; i < len(quantities); i++ {
		scaled := quantities[i] * float64(multiplier)

		// fmt.Printf("q: %.f, p: %d * %.2f = %.f\n", quantities[i], portions, multiplier, scaled)
		scaledQuantities[i] = scaled
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
