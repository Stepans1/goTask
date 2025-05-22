package generationService

import (
	"errors"
	"fmt"
	"goTask/internal/DTO"
	"goTask/internal/constants"
	"math/rand"
	"slices"
)

var previousPassword string

type PasswordGenerationService interface {
	Generate(options DTO.GenerationOptions) (string, []string)
	GetGenerationOptions() []DTO.PaternDescription
}

type passwordGenerationService struct{}

func NewPasswordGenerationService() PasswordGenerationService {
	return &passwordGenerationService{}
}

func (s *passwordGenerationService) Generate(options DTO.GenerationOptions) (string, []string) {
	length := options.Length
	selectedSets := options.Options

	validationErrors := validateGenerationOptions(length, selectedSets)
	if len(validationErrors) > 0 {
		messages := make([]string, len(validationErrors))
		for index, err := range validationErrors {
			messages[index] = err.Error()
		}

		return "", messages
	}

	charSetMap := buildCharSetMap()
	availableChars := []byte{}
	for _, set := range selectedSets {
		availableChars = append(availableChars, charSetMap[set]...)
	}

	if len(availableChars) < length {
		return "", []string{"not enough unique characters to generate a password of length"}
	}

	const maxAttempts = 5
	for range maxAttempts {
		password := generatePassword(length, selectedSets, availableChars, charSetMap)
		if password != previousPassword {
			previousPassword = password
			return password, nil
		}
	}

	return "", []string{"failed to generate a unique password after several attempts"}
}

func (s *passwordGenerationService) GetGenerationOptions() []DTO.PaternDescription {
	return constants.GetPatternsDescriptions()
}

func generatePassword(length int, selectedSets []string, availableChars []byte, charSetMap map[string][]byte) string {
	mandatoryCount := min(len(selectedSets), length)
	mandatoryChars := []byte{}
	used := make(map[byte]bool)

	for i := range mandatoryCount {
		set := charSetMap[selectedSets[i]]
		ch := set[rand.Intn(len(set))]
		mandatoryChars = append(mandatoryChars, ch)
		used[ch] = true
	}

	remainingPool := []byte{}
	for _, ch := range availableChars {
		if !used[ch] {
			remainingPool = append(remainingPool, ch)
		}
	}

	remainingLen := length - len(mandatoryChars)
	for i := 0; i < remainingLen && len(remainingPool) > 0; i++ {
		idx := rand.Intn(len(remainingPool))
		ch := remainingPool[idx]
		mandatoryChars = append(mandatoryChars, ch)
		remainingPool = slices.Delete(remainingPool, idx, idx+1)
	}

	rand.Shuffle(len(mandatoryChars), func(i, j int) {
		mandatoryChars[i], mandatoryChars[j] = mandatoryChars[j], mandatoryChars[i]
	})

	return string(mandatoryChars)
}

func validateGenerationOptions(length int, options []string) []error {
	var errs []error

	if length < 1 {
		errs = append(errs, errors.New("the password length must be at least 1"))
	}

	if len(options) == 0 {
		errs = append(errs, errors.New("you must select at least one character set"))
	}

	for _, opt := range options {
		if !isAllowedOption(opt) {
			errs = append(errs, fmt.Errorf("invalid option: %s", opt))
		}
	}

	if length < len(options) {
		errs = append(errs, errors.New("the password length cannot be less than the number of selected character sets"))
	}

	return errs
}

func isAllowedOption(opt string) bool {
	return slices.Contains(constants.GetAllowedOptions(), opt)
}

func buildCharSetMap() map[string][]byte {
	chMap := make(map[string][]byte)
	for _, cs := range constants.GetCharSets() {
		chMap[cs.Key] = []byte(cs.Set)
	}

	return chMap
}
