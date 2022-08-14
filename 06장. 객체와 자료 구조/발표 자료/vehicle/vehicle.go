package vehicle

// Vehicle1 보다 Vehicle2가 더 좋은 것 같다.

type Vehicle1 interface {
	FuelTankCapacityInGallons() int
	GallonsOfGasoline() int
}

type Vehicle2 interface {
	GetPercentFuelRemaining() int
}
