package validator

type Case struct {
	Hash     string
	Password string
	Solved   bool
}

type Validator struct {
	cases      []Case
	answerChan chan [32]byte
}

func NewValidator(answerChan chan [32]byte) *Validator {
	return &Validator{
		answerChan: answerChan,
	}
}

// Start validating answers
func (v *Validator) Start() []Case {
	solved := 0

	for a := range v.answerChan {
		// loop end condition
		if solved == len(v.cases) {
			return v.cases
		}

		// case validation
		for i, c := range v.cases {
			if c.Solved {
				continue
			}
			if c.Hash == string(a[:]) {
				v.cases[i].Solved = true
				v.cases[i].Password = string(a[:])
				solved++
			}
		}
	}

	return []Case{}
}
