package internal

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportId     string
	CountryId      string
}

func (p *Passport) AddInformation(info string, validate bool) {
	parts := strings.Split(info, " ")

	for _, part := range parts {
		keyVal := strings.Split(part, ":")

		switch keyVal[0] {
		case "byr":
			if validate {
				p.BirthYear = parseBirthYear(keyVal[1])
			} else {
				p.BirthYear = keyVal[1]
			}
		case "iyr":
			if validate {
				p.IssueYear = parseIssueYear(keyVal[1])
			} else {
				p.IssueYear = keyVal[1]
			}
		case "eyr":
			if validate {
				p.ExpirationYear = parseExpirationYear(keyVal[1])
			} else {
				p.ExpirationYear = keyVal[1]
			}
		case "hgt":
			if validate {
				p.Height = parseHeight(keyVal[1])
			} else {
				p.Height = keyVal[1]
			}
		case "hcl":
			if validate {
				p.HairColor = parseHairColor(keyVal[1])
			} else {
				p.HairColor = keyVal[1]
			}
		case "ecl":
			if validate {
				p.EyeColor = parseEyeColor(keyVal[1])
			} else {
				p.EyeColor = keyVal[1]
			}
		case "pid":
			if validate {
				p.PassportId = parsePassportId(keyVal[1])
			} else {
				p.PassportId = keyVal[1]
			}
		case "cid":
			p.CountryId = keyVal[1]
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

func parseBirthYear(birthYear string) string {
	year, err := strconv.Atoi(birthYear)
	if err != nil {
		return ""
	}

	if 1920 <= year && year <= 2002 {
		return birthYear
	}

	return ""
}

func parseIssueYear(issueYear string) string {
	year, err := strconv.Atoi(issueYear)
	if err != nil {
		return ""
	}

	if 2010 <= year && year <= 2020 {
		return issueYear
	}

	return ""
}

func parseExpirationYear(expYear string) string {
	year, err := strconv.Atoi(expYear)
	if err != nil {
		return ""
	}

	if 2020 <= year && year <= 2030 {
		return expYear
	}

	return ""
}

func parseHeight(height string) string {
	regex, err := regexp.Compile("^([1-9][0-9][0-9]?)(cm|in)$")
	if err != nil {
		log.Fatal("Issue with regex for parse height")
	}

	parts := regex.FindStringSubmatch(height)
	if len(parts) == 0 {
		return ""
	}

	if parts[2] == "cm" {
		h, err := strconv.Atoi(parts[1])
		if err != nil {
			return ""
		}

		if 150 <= h && h <= 193 {
			return height
		}

		return ""
	} else if parts[2] == "in" {
		h, err := strconv.Atoi(parts[1])
		if err != nil {
			return ""
		}

		if 59 <= h && h <= 76 {
			return height
		}
		return ""
	}

	return ""
}

func parseHairColor(hairColor string) string {
	matched, err := regexp.MatchString("^#[0-9a-f]{6}$", hairColor)
	if err != nil || !matched {
		return ""
	}

	return hairColor
}

func parseEyeColor(eyeColor string) string {
	if eyeColor == "amb" ||
		eyeColor == "blu" ||
		eyeColor == "brn" ||
		eyeColor == "gry" ||
		eyeColor == "grn" ||
		eyeColor == "hzl" ||
		eyeColor == "oth" {
		return eyeColor
	}

	return ""
}

func parsePassportId(passportId string) string {
	matched, err := regexp.MatchString("^[0-9]{9}$", passportId)
	if err != nil || !matched {
		return ""
	}

	return passportId
}
