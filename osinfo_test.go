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
	osRelease := `NAME="Alpine Linux"
ID=alpine
VERSION_ID=3.8.0
PRETTY_NAME="Alpine Linux v3.8"
HOME_URL="http://alpinelinux.org"
BUG_REPORT_URL="http://bugs.alpinelinux.org"
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	// Alpine has no /etc/lsb-release

	expectEqualStrings(t, "alpine", info.ID)
	expectEqualStrings(t, "3.8.0", info.Version)
	expectEqualStrings(t, "Alpine Linux", info.Name)
	expectEqualStrings(t, "", info.Codename)
}

func TestCentos(t *testing.T) {
	osRelease := `NAME="CentOS Linux"
VERSION="8 (Core)"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="8"
PLATFORM_ID="platform:el8"
PRETTY_NAME="CentOS Linux 8 (Core)"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:8"
HOME_URL="https://www.centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"

CENTOS_MANTISBT_PROJECT="CentOS-8"
CENTOS_MANTISBT_PROJECT_VERSION="8"
REDHAT_SUPPORT_PRODUCT="centos"
REDHAT_SUPPORT_PRODUCT_VERSION="8"

`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	// CentOS has no /etc/lsb-release

	expectEqualStrings(t, "centos", info.ID)
	expectEqualStrings(t, "8", info.Version)
	expectEqualStrings(t, "CentOS Linux", info.Name)
	expectEqualStrings(t, "", info.Codename)
}

func TestDebian(t *testing.T) {
	osRelease := `PRETTY_NAME="Debian GNU/Linux 9 (stretch)"
NAME="Debian GNU/Linux"
VERSION_ID="9"
VERSION="9 (stretch)"
VERSION_CODENAME=stretch
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	// Debian has no /etc/lsb-release

	expectEqualStrings(t, "debian", info.ID)
	expectEqualStrings(t, "9", info.Version)
	expectEqualStrings(t, "Debian GNU/Linux", info.Name)
	expectEqualStrings(t, "stretch", info.Codename)
}

func TestFedora(t *testing.T) {
	osRelease := `NAME=Fedora
VERSION="31 (Container Image)"
ID=fedora
VERSION_ID=31
VERSION_CODENAME=""
PLATFORM_ID="platform:f31"
PRETTY_NAME="Fedora 31 (Container Image)"
ANSI_COLOR="0;34"
LOGO=fedora-logo-icon
CPE_NAME="cpe:/o:fedoraproject:fedora:31"
HOME_URL="https://fedoraproject.org/"
DOCUMENTATION_URL="https://docs.fedoraproject.org/en-US/fedora/f31/system-administrators-guide/"
SUPPORT_URL="https://fedoraproject.org/wiki/Communicating_and_getting_help"
BUG_REPORT_URL="https://bugzilla.redhat.com/"
REDHAT_BUGZILLA_PRODUCT="Fedora"
REDHAT_BUGZILLA_PRODUCT_VERSION=31
REDHAT_SUPPORT_PRODUCT="Fedora"
REDHAT_SUPPORT_PRODUCT_VERSION=31
PRIVACY_POLICY_URL="https://fedoraproject.org/wiki/Legal:PrivacyPolicy"
VARIANT="Container Image"
VARIANT_ID=container
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	// Fedora has no /etc/lsb-release

	expectEqualStrings(t, "fedora", info.ID)
	expectEqualStrings(t, "31", info.Version)
	expectEqualStrings(t, "Fedora", info.Name)
	expectEqualStrings(t, "", info.Codename)
}

func TestGentoo(t *testing.T) {
	osRelease := `NAME=Gentoo
ID=gentoo
PRETTY_NAME="Gentoo/Linux"
ANSI_COLOR="1;32"
HOME_URL="https://www.gentoo.org/"
SUPPORT_URL="https://www.gentoo.org/support/"
BUG_REPORT_URL="https://bugs.gentoo.org/"
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	// Gentoo has no /etc/lsb-release

	expectEqualStrings(t, "gentoo", info.ID)
	expectEqualStrings(t, "", info.Version)
	expectEqualStrings(t, "Gentoo", info.Name)
	expectEqualStrings(t, "", info.Codename)
}

func TestKali(t *testing.T) {
	osRelease := `PRETTY_NAME="Kali GNU/Linux Rolling"
NAME="Kali GNU/Linux"
ID=kali
VERSION="2020.2"
VERSION_ID="2020.2"
VERSION_CODENAME="kali-rolling"
ID_LIKE=debian
ANSI_COLOR="1;31"
HOME_URL="https://www.kali.org/"
SUPPORT_URL="https://forums.kali.org/"
BUG_REPORT_URL="https://bugs.kali.org/"
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	// Kali has no /etc/lsb-release

	expectEqualStrings(t, "kali", info.ID)
	expectEqualStrings(t, "2020.2", info.Version)
	expectEqualStrings(t, "Kali GNU/Linux", info.Name)
	expectEqualStrings(t, "kali-rolling", info.Codename)
}

func TestManjaro(t *testing.T) {
	osRelease := `NAME="Manjaro Linux"
ID=manjaro
ID_LIKE=arch
PRETTY_NAME="Manjaro Linux"
ANSI_COLOR="1;32"
HOME_URL="https://www.manjaro.org/"
SUPPORT_URL="https://www.manjaro.org/"
BUG_REPORT_URL="https://bugs.manjaro.org/"
LOGO=manjarolinux
`
	lsbRelease := `DISTRIB_ID=ManjaroLinux
DISTRIB_RELEASE=19.0.2
DISTRIB_CODENAME=Kyria
DISTRIB_DESCRIPTION="Manjaro Linux"
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	parseEtcLSBRelease(info, lsbRelease)

	expectEqualStrings(t, "manjaro", info.ID)
	expectEqualStrings(t, "19.0.2", info.Version)
	expectEqualStrings(t, "Manjaro Linux", info.Name)
	expectEqualStrings(t, "Kyria", info.Codename)
}

func TestOpenSUSE(t *testing.T) {
	osRelease := `NAME="openSUSE Leap"
VERSION="15.1"
ID="opensuse-leap"
ID_LIKE="suse opensuse"
VERSION_ID="15.1"
PRETTY_NAME="openSUSE Leap 15.1"
ANSI_COLOR="0;32"
CPE_NAME="cpe:/o:opensuse:leap:15.1"
BUG_REPORT_URL="https://bugs.opensuse.org"
HOME_URL="https://www.opensuse.org/"
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	// OpenSUSE has no /etc/lsb-release

	expectEqualStrings(t, "opensuse-leap", info.ID)
	expectEqualStrings(t, "15.1", info.Version)
	expectEqualStrings(t, "openSUSE Leap", info.Name)
	expectEqualStrings(t, "", info.Codename)
}

func TestOracle(t *testing.T) {
	osRelease := `NAME="Oracle Linux Server"
VERSION="8.1"
ID="ol"
ID_LIKE="fedora"
VARIANT="Server"
VARIANT_ID="server"
VERSION_ID="8.1"
PLATFORM_ID="platform:el8"
PRETTY_NAME="Oracle Linux Server 8.1"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:oracle:linux:8:1:server"
HOME_URL="https://linux.oracle.com/"
BUG_REPORT_URL="https://bugzilla.oracle.com/"

ORACLE_BUGZILLA_PRODUCT="Oracle Linux 8"
ORACLE_BUGZILLA_PRODUCT_VERSION=8.1
ORACLE_SUPPORT_PRODUCT="Oracle Linux"
ORACLE_SUPPORT_PRODUCT_VERSION=8.1
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	// Oracle has no /etc/lsb-release

	expectEqualStrings(t, "ol", info.ID)
	expectEqualStrings(t, "8.1", info.Version)
	expectEqualStrings(t, "Oracle Linux Server", info.Name)
	expectEqualStrings(t, "", info.Codename)
}

func TestUbuntu(t *testing.T) {
	osRelease := `NAME="Ubuntu"
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
	lsbRelease := `DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=19.10
DISTRIB_CODENAME=eoan
DISTRIB_DESCRIPTION="Ubuntu 19.10"
`

	info := new(OSInfo)
	parseEtcOSRelease(info, osRelease)
	parseEtcLSBRelease(info, lsbRelease)

	expectEqualStrings(t, "ubuntu", info.ID)
	expectEqualStrings(t, "19.10", info.Version)
	expectEqualStrings(t, "Ubuntu", info.Name)
	expectEqualStrings(t, "eoan", info.Codename)
}

func TestMacOSSierra(t *testing.T) {
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

func TestWSL(t *testing.T) {
	osRelease := `NAME="Ubuntu"
VERSION="20.04 LTS (Focal Fossa)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 20.04 LTS"
VERSION_ID="20.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=focal
UBUNTU_CODENAME=focal`

	info := new(OSInfo)
	info.IsWSL = true // Simulate WSL detection
	parseEtcOSRelease(info, osRelease)

	expectEqualStrings(t, "ubuntu", info.ID)
	expectEqualStrings(t, "20.04", info.Version)
	expectEqualStrings(t, "Ubuntu (WSL)", info.Name)
	expectEqualStrings(t, "focal", info.Codename)
	if !info.IsWSL {
		t.Error("Expected IsWSL to be true")
	}
}
