OS Info
=======

This package provides a cross-platform way to identify the hardware your go code is running on.

The following fields are provided by the `OSInfo` struct:

| Field        | Description                             |
| ------------ | --------------------------------------- |
| Family       | The OS type as defined by `GOOS`        |
| Architecture | The architecture as defined by `GOARCH` |
| ID           | The OS ID as defined by the OS          |
| Name         | The OS name as defined by the OS        |
| Codename     | The release codename (if any)           |
| Version      | The release version                     |
| Build        | The build number (if any)               |

Supported Operating Systems
---------------------------

The following operating systems are currently supported:

- Linux
- FreeBSD
- macOS
- Windows

If you wish to see another operating system supported (provided it is in
[this list](https://github.com/golang/go/blob/master/src/go/build/syslist.go)),
please open a pull request with the necessary changes and tests.
Use one of the existing `getOSInfoXYZ()` functions as a guide (most commonly
it involves parsing the output from a command or file).

Usage
-----

```golang
	info, err := osinfo.GetOSInfo()
	if err != nil {
		// TODO: Handle this
	}

	fmt.Printf("Family:       %v\n", info.Family)
	fmt.Printf("Architecture: %v\n", info.Architecture)
	fmt.Printf("ID:           %v\n", info.ID)
	fmt.Printf("Name:         %v\n", info.Name)
	fmt.Printf("Codename:     %v\n", info.Codename)
	fmt.Printf("Version:      %v\n", info.Version)
	fmt.Printf("Build:        %v\n", info.Build)
```

### Output on various platforms

#### Ubuntu Linux

```
Family:       linux
Architecture: amd64
ID:           ubuntu
Name:         Ubuntu
Codename:     eoan
Version:      19.10
Build:
```

### Alpine Linux

```
Family:       linux
Architecture: amd64
ID:           alpine
Name:         Alpine Linux
Codename:
Version:      3.8.0
Build:
```

#### Windows

```
Family:       windows
Architecture: amd64
ID:           windows
Name:         Windows 10 Pro
Codename:     1903
Version:      10.0
Build:        18362
```

#### FreeBSD

```
Family:       freebsd
Architecture: amd64
ID:           freebsd
Name:         FreeBSD
Codename:
Version:      12.0-RELEASE
Build:        r341666
```

#### Mac OS

```
Family:       darwin
Architecture: amd64
ID:           darwin
Name:         Mac OS X
Codename:     Sierra
Version:      10.12.6
Build:        16G2136
```

License
-------

Copyright (c) 2019 Blackfire SAS (https://blackfire.io). All rights reserved.

License type: MIT

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial
portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
