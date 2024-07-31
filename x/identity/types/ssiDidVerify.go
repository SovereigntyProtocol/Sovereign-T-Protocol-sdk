package types

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

var networkNames = map[string]string{
	"1":  "discoverynet",
	"2":  "identityforge",
	"3":  "questchain",
	"4":  "authenticitynet",
	"5":  "personapulse",
	"6":  "veritysphere",
	"7":  "insightnet",
	"8":  "emergenet",
	"9":  "egonet",
	"10": "identitynexa",
}

const AccAddprifix = "ssi"

func GetDefaultDidPrefix(networkID string) string {
	networkName, exists := networkNames[networkID]
	if !exists {
		networkName = "UnknownNetwork" // or handle the error as needed
	}
	return "did:sovid:" + networkName + ":"
}

// findDividerInDid finds the divider "1" in the string and return all characters before "1"
func findPrefixAndDataInDid(s string) (string, string, error) {
	for i, ch := range s {
		if ch == '1' {
			if len(s[i:]) == 0 {
				return s[:i], "", nil
			}
			return s[:i], s[i:], nil
		}
	}
	return "", "", errors.New("divider not found in the string")
}

func GetPrefixFromDidString(did string) (string, error) {
	if len(did) == 0 {
		return "", fmt.Errorf("empty did string is not allowed")
	}

	prifix, _, err := findPrefixAndDataInDid(did)

	if err != nil {
		return "", err
	}
	return prifix, nil
}

func GetDataAfterDivider(did string) (string, error) {

	if len(did) == 0 {
		return "", fmt.Errorf("empty did string is not allowed")
	}
	_, data, err := findPrefixAndDataInDid(did)

	if err != nil {
		return "", err
	}
	// for i, ch := range did {
	// 	if ch == '1' {
	// 		if len(did[i:]) == 0 {
	// 			return "", errors.New("data does not exist")
	// 		}
	// 		return did[i:], nil
	// 	}
	// }
	return data, nil
}

func VerifyPrefixFormat(did string) (string, error) {

	didprifix := GetDefaultDidPrefix("4")

	if len(strings.TrimSpace(did)) == 0 {
		return "", fmt.Errorf("empty did string is not allowed")
	}

	fountprefix, err := GetPrefixFromDidString(did)

	if err != nil {
		return "", err
	}

	if fountprefix != didprifix {
		return "", fmt.Errorf("invalid did prefix; expected %s, got %s", didprifix, fountprefix)
	}

	return fountprefix, nil
}

func VerifyDidData(did string) (string, error) {

	data, detaErr := GetDataAfterDivider(did)

	if detaErr != nil {
		return "", detaErr
	}
	return data, nil
}

func VerifyDidFormat(did string) (bool, error) {

	_, Prefixerr := VerifyPrefixFormat(did)
	if Prefixerr != nil {
		return false, Prefixerr
	}

	data, Dataerr := VerifyDidData(did)
	if Dataerr != nil {
		return false, Dataerr
	}

	_, _, err := bech32.DecodeAndConvert(AccAddprifix + data)
	if err != nil {
		return false, err
	}

	return true, nil
}
