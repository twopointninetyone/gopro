package main

type Items struct {
	Name string
	Value int
}

type Research struct {
	Name string
	RPcost int
	RequiredResearches map[string]*ResearchRequirements
}

type ResearchRequirements struct {
	RequiredResearches map[string]*Research
}

type ResearchTree struct {
	Researches map[string]*Research
}

type Machine struct {
	Name string
	Cost int
	MachineType string
	ItemInput map[Items]any
	ItemYield map[Items]any
	PowerInput float32
	PowerOutput float32 /// for generators
	ResearchNeeded string
}

type MachineList struct {
	ListName string
	Tier int
	Machines map[string]*Machine
}


