package main

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main(){
//	ExampleParsePKIXPublicKey()
	ExampleCertificate_Verify()
}

func ExampleCertificate_Verify() {
	// Verifying with a custom list of root certificates.

	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIDgzCCAmugAwIBAgIJAJYAgNGYyzWhMA0GCSqGSIb3DQEBBQUAMFcxCzAJBgNV
BAYTAkZJMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxDjAMBgNVBAoMBU5PS0lBMQww
CgYDVQQLDANPU1MxEzARBgNVBAMMCk5FM1MgV1MgQ0EwIBcNMTUwOTI4MDk0NTU2
WhgPMjExNTA5MDQwOTQ1NTZaMFcxCzAJBgNVBAYTAkZJMRUwEwYDVQQHDAxEZWZh
dWx0IENpdHkxDjAMBgNVBAoMBU5PS0lBMQwwCgYDVQQLDANPU1MxEzARBgNVBAMM
Ck5FM1MgV1MgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDLBtUF
oyHGJMVt1TaiZ9OPNL2+9rFRRBIMtN/u7bDnQEhRTGfUc0AB0XlzHOZFO9/O2Nb3
FIcV+lEQUyIioQEmrGb9eHEbhYrS/PzA31YI8UFpkN8EPuZgCjc1ZpLCz1cZAryk
5z32fzMMAlQtxW9owAUzz0VuzCNwuqClhfbW2a4wpJwJ59URDKKV0pRO2lR4t1Vq
GU8bwA8iCt0XPmHm6MiHCJcRSughwmfsoaOx07BK5cG0vqO4B0bJbSzeEcwMc+Go
SKoJYiHlzreOE/wposMdCtGdwOMZz5ORXV0ptC5ehinKakO1W1MT8OmnC2SzzvvX
SiqvoObzFVpjsPOfAgMBAAGjUDBOMB0GA1UdDgQWBBQ4jQDFYbhFeFGjBt0pQgVA
FBJCLDAfBgNVHSMEGDAWgBQ4jQDFYbhFeFGjBt0pQgVAFBJCLDAMBgNVHRMEBTAD
AQH/MA0GCSqGSIb3DQEBBQUAA4IBAQAyWMYtqkiHbr2zoAHdEry2U4lAiJmZdgdx
d72hKCLuVHzFHSlrfP11jjRcbp0i3oILV1txByp/AH/NkwLTXrdGGhR2/te/PJPO
lxcwDCHW/g/Ypdz1obYD7Xb/67XWryL7QTVJ4oU6tYUbiE6coynVzEJg1gHkJ53D
B+5eZX3GOCQDfEPPeRtH5qAoZ3uws5n8q8MWCDUEd0wLkKNmutVOoQALvoeQbSO+
wV6j9g58VDXK+gNlZVH/StA6VnKxZm9jSBA2DFZcq/3MGH7hut4eF4tjEfE/x2jP
k+ideUuT2WpY2blg2bIm0bgh3Of11l355C1kB3k9S8lw343Ki8fd
-----END CERTIFICATE-----`

	const certPEM = `
-----BEGIN CERTIFICATE-----
MIIDODCCAiCgAwIBAgICEAAwDQYJKoZIhvcNAQEFBQAwVzELMAkGA1UEBhMCRkkx
FTATBgNVBAcMDERlZmF1bHQgQ2l0eTEOMAwGA1UECgwFTk9LSUExDDAKBgNVBAsM
A09TUzETMBEGA1UEAwwKTkUzUyBXUyBDQTAgFw0xNTA5MjgxMTE0NTRaGA8yMTE1
MDkwNDExMTQ1NFowbDELMAkGA1UEBhMCRkkxEjAQBgNVBAgTCUJlcmtzaGlyZTEQ
MA4GA1UEBxMHTmV3YnVyeTEOMAwGA1UEChMFTm9raWExDDAKBgNVBAsTA09TUzEZ
MBcGA1UEAxMQTkUzU1dTIE9NQSBBZ2VudDCBnzANBgkqhkiG9w0BAQEFAAOBjQAw
gYkCgYEAv69FC4JjrefgV+H7GbhnQFxKMZG5msXLCmFIPmEPkySchF2CAN8Iz7vy
upo1Y77XH9lVX2ELSHC0p4xcfDKwAaZIvk3QtzVEq9tZ1fm9F+Sps4RkeP7RZzoA
cKFxLW3BUL+6yd6jppvgMf5X1Nyn7msJ5ZsuAqIA5KyDGak64a0CAwEAAaN7MHkw
CQYDVR0TBAIwADAsBglghkgBhvhCAQ0EHxYdT3BlblNTTCBHZW5lcmF0ZWQgQ2Vy
dGlmaWNhdGUwHQYDVR0OBBYEFMo+2VseArKRdm/Wa7r3psjc9P6dMB8GA1UdIwQY
MBaAFDiNAMVhuEV4UaMG3SlCBUAUEkIsMA0GCSqGSIb3DQEBBQUAA4IBAQA0dmnQ
cwNBoR2njUv4AdumXhdjk+5zdMUPF14rUUnm7v3n28l00mnOcFSPZ4JZcLpzmfhR
XetEemiF9pH+20wctvMuzBVWPA8EMtkgfboA6PbepfDE1cxTxESHaQpB4Yq8E5cG
1pppEi1/jgoRVytpKiCIKFSLVJERgVktS/8U4y0gFzwSXuKbgu91S4coXKsi5eA2
qzNcehZmjua7IoKvMdzSD76JBdtZEn2aegdqNR0Wka/gbrytzZ3G07JKcXhseql4
pj34xpQkfoTqKfnGsbUfxqYk3ZTtILduP/V6wZFCJF+iHNk9kPZzD4PUMGt5FuLM
nL9vHp5iW3RlIDHV
-----END CERTIFICATE-----`

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		Roots:   roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}
	fmt.Println("SUCCESS")
}

func ExampleParsePKIXPublicKey() {
	const pubPEM = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1331XRMk5XJShvSTljb6
8MRBjtsgHbhTTx5bF93NFN2IoSduRUVsAWxx5jwrvRXxmm6fus2neK0PIYqEXuz9
G3AGx8uu1ASnB8+IhY6PbrY+YSnhJV4TXphVwmOhobeigH1XiF/Q6O08342h1xP7
4SbqrQapjYvsLUclyvNJB7jX1+n81samVqPFZ03z0GxAGnpZsPpGG997OKiy++Eu
OhH0zobVTV4PGLHAvTSDJu/CBvaKjSUPGqkcVatEZo6eVYhJWaQDetxg9dsS+dne
pmr+5eduoCjccwyPOG+4c53ZJPjUFCI2mQvslVqqEM27KEqldG6J1wjMjAtL/Y3T
QQIDAQAB
-----END PUBLIC KEY-----`

	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}
	fmt.Println(*block)

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		fmt.Println("pub is of type RSA:", *pub)
	case *dsa.PublicKey:
		fmt.Println("pub is of type DSA:", pub)
	case *ecdsa.PublicKey:
		fmt.Println("pub is of type ECDSA:", pub)
	default:
		panic("unknown type of public key")
	}
}

