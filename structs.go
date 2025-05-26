package main

type Items struct {
	Name string
	Value int
}

type Research struct {
	Name string
	RPcost int
	RequiredResearches map[*ResearchRequirements]any
}

type ResearchRequirements struct {
	RequiredResearches ResearchTree /// basically tree of requirements needed
}

type ResearchTree struct {
	Researches map[*Research]any
}

type Machine struct {
	Name string
	Cost int
	MachineType string
	ItemInput map[Items]any
	ItemYield map[Items]any
	PowerInput float32
	PowerOutput float32 /// for generators
	ResearchNeeded Research
}

type MachineList struct {
	ListName string
	Tier int
	Machines map[string]*Machine
}




