package osinfo

import (
	"fmt"
	"testing"
)

func expectEqualStrings(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected [%v] but got [%v]", expected, actual)
	}
}

func expectEqualInts(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected [%v] but got [%v]", expected, actual)
	}
}

func TestAlpine(t *testing.T) {
	testData := `NAME="Alpine Linux"
ID=alpine
VERSION_ID=3.8.0
PRETTY_NAME="Alpine Linux v3.8"
HOME_URL="http://alpinelinux.org"
BUG_REPORT_URL="http://bugs.alpinelinux.org"
`

	info := new(OSInfo)
	err := parseEtcOSRelease(info, testData)
	if err != nil {
		t.Error(err)
	}

	expectEqualStrings(t, "alpine", info.ID)
	expectEqualStrings(t, "3.8.0", info.Version)
	expectEqualStrings(t, "Alpine Linux", info.Name)
}

func TestUbuntu(t *testing.T) {
	testData := `NAME="Ubuntu"
VERSION="19.10 (Eoan Ermine)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 19.10"
VERSION_ID="19.10"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=eoan
UBUNTU_CODENAME=eoan
`

	info := new(OSInfo)
	err := parseEtcOSRelease(info, testData)
	if err != nil {
		t.Error(err)
	}

	expectEqualStrings(t, "ubuntu", info.ID)
	expectEqualStrings(t, "19.10", info.Version)
	expectEqualStrings(t, "Ubuntu", info.Name)
	expectEqualStrings(t, "eoan", info.Codename)
}

func TestSierra(t *testing.T) {
	info := new(OSInfo)
	productVersion := "10.12.6"
	buildVersion := "16G1815"
	err := parseMacSWVers(info, productVersion, buildVersion)
	if err != nil {
		t.Error(err)
	}

	expectEqualStrings(t, "10.12.6", info.Version)
	expectEqualStrings(t, "Sierra", info.Codename)
	expectEqualStrings(t, "16G1815", info.Build)
}

func TestFreeBSD(t *testing.T) {
	info := new(OSInfo)
	unameV := "FreeBSD 12.0-RELEASE r341666 GENERIC"
	err := parseFreeBSDUname(info, unameV)
	if err != nil {
		t.Error(err)
	}

	expectEqualStrings(t, "12.0-RELEASE", info.Version)
	expectEqualStrings(t, "r341666", info.Build)
	expectEqualStrings(t, "FreeBSD", info.Name)
}

func expectRegistryString(t *testing.T, expected string, id string, regOutput string) {
	result, err := extractRegistryString(id, regOutput)
	if err != nil {
		t.Error(err)
	}
	expectEqualStrings(t, expected, result)
}

func expectRegistryInt(t *testing.T, expected int, id string, regOutput string) {
	result, err := extractRegistryInt(id, regOutput)
	if err != nil {
		t.Error(err)
	}
	expectEqualInts(t, expected, result)
}

func TestWindowsExtractProductName(t *testing.T) {
	expectRegistryString(t, "Windows 10 Pro", "ProductName", `
HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion
    ProductName    REG_SZ    Windows 10 Pro
`)
}

func TestWindowsExtractCurrentVersion(t *testing.T) {
	expectRegistryString(t, "6.3", "CurrentVersion", `
HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion
    CurrentVersion    REG_SZ    6.3
`)
}

func TestWindowsExtractMajorVersion(t *testing.T) {
	expectRegistryInt(t, 10, "CurrentMajorVersionNumber", `
HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion
    CurrentMajorVersionNumber    REG_DWORD    0xa
`)
}

func TestWindowsExtractMinorVersion(t *testing.T) {
	expectRegistryInt(t, 0, "CurrentMinorVersionNumber", `
HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion
    CurrentMinorVersionNumber    REG_DWORD    0x0
`)
}

func TestWindowsExtractCurrentBuild(t *testing.T) {
	expectRegistryString(t, "18362", "CurrentBuild", `
HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion
    CurrentBuild    REG_SZ    18362
`)
}

func TestDemonstrate(t *testing.T) {
	info, err := GetOSInfo()
	if err != nil {
		t.Errorf("Error while getting OS info: %v", err)
	}

	fmt.Printf("Family:       %v\n", info.Family)
	fmt.Printf("Architecture: %v\n", info.Architecture)
	fmt.Printf("ID:           %v\n", info.ID)
	fmt.Printf("Name:         %v\n", info.Name)
	fmt.Printf("Codename:     %v\n", info.Codename)
	fmt.Printf("Version:      %v\n", info.Version)
	fmt.Printf("Build:        %v\n", info.Build)
}
