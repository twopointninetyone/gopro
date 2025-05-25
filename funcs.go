package main

/// Research Functions

func createTree() *ResearchTree {
	return &ResearchTree {
		Researches: make(map[string]*Research),
	}
}

func (rt *ResearchTree) AddResearch(r *Research) {
	if rt.Researches == nil {
		rt.Researches = make(map[string]*Research)
	}
	rt.Researches[r.Name] = r
}

/// Machine Functions

func createMachineList(lstName string, tier int) *MachineList {
	return &MachineList {
		ListName: lstName,
		Tier: tier,
	}
}

func (lst *MachineList) addMachine(m *Machine) {
	if lst.Machines == nil {
		lst.Machines = make(map[string]*Machine)
	}
	lst.Machines[m.Name] = m
}
