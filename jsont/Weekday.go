package jsont

type Weekday int

const (
	//constant generator iota, which is used to create a sequence 
	//of related values without spelling out each one explicitly. 
	//In a const declaration, the value of iota begins at zero and 
	//increments by one for each item in the sequence.
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
