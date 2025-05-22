package constants

import "goTask/internal/DTO"

const (
	NumbersKey      = "numbers"
	SmallLettersKey = "smallLetters"
	BigLettersKey   = "bigLetters"
)

const (
	numbersDescription      = "Numbers (0-9)"
	smallLettersDescription = "Lowercase letters (a-z)"
	bigLettersDescription   = "Uppercase letters (A-Z)"
)

const (
	numbersSet      = "0123456789"
	smallLettersSet = "abcdefghijklmnopqrstuvwxyz"
	bigLettersSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetCharSets() []DTO.CharSet {
	return []DTO.CharSet {
		{
			Key: NumbersKey,
			Set: numbersSet,
		},
		{
			Key: SmallLettersKey,
			Set: smallLettersSet,
		},
		{
			Key: BigLettersKey,
			Set: bigLettersSet,
		},
	}
}

func GetAllowedOptions() []string {
	return []string{
		NumbersKey,
		SmallLettersKey,
		BigLettersKey,
	}
}

func GetPatternsDescriptions() []DTO.PaternDescription {
	return []DTO.PaternDescription{
		{
			Key:         NumbersKey,
			Description: numbersDescription,
		},
		{
			Key:         SmallLettersKey,
			Description: smallLettersDescription,
		},
		{
			Key:         BigLettersKey,
			Description: bigLettersDescription,
		},
	}
}
