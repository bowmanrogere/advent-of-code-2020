package internal

import "strings"

type Passport struct {
	BirthYear string
	IssueYear string
	ExpirationYear string
	Height string
	HairColor string
	EyeColor string
	PassportId string
	CountryId string
}

func (p *Passport) AddInformation(info string) {
	parts := strings.Split(info, " ")

	for _, part := range parts {
		keyVal := strings.Split(part, ":")

		switch keyVal[0] {
		case "byr":
			p.BirthYear = keyVal[1]
		case "iyr":
			p.IssueYear = keyVal[1]
		case "eyr":
			p.ExpirationYear = keyVal[1]
		case "hgt":
			p.Height = keyVal[1]
		case "hcl":
			p.HairColor = keyVal[1]
		case "ecl":
			p.EyeColor = keyVal[1]
		case "pid":
			p.PassportId = keyVal[1]
		case "cid":
			p.CountryId	= keyVal[1]
		}
	}
}

func (p *Passport) Valid() bool {
	return p.BirthYear != "" &&
		p.IssueYear != "" &&
		p.ExpirationYear != "" &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.PassportId != ""
}
