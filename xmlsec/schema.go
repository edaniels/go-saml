package xmlsec

import (
	"encoding/xml"
)

// Method is part of Signature.
type Method struct {
	Algorithm string `xml:",attr"`
}

// Signature is a model for the Signature object specified by XMLDSIG. This is
// convenience object when constructing XML that you'd like to sign. For example:
//
//    type Foo struct {
//       Stuff string
//       Signature Signature
//    }
//
//    f := Foo{Suff: "hello"}
//    f.Signature = DefaultSignature()
//    buf, _ := xml.Marshal(f)
//    buf, _ = Sign(key, buf)
//
type Signature struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`

	CanonicalizationMethod Method             `xml:"SignedInfo>CanonicalizationMethod"`
	SignatureMethod        Method             `xml:"SignedInfo>SignatureMethod"`
	ReferenceTransforms    []Method           `xml:"SignedInfo>Reference>Transforms>Transform"`
	DigestMethod           Method             `xml:"SignedInfo>Reference>DigestMethod"`
	DigestValue            string             `xml:"SignedInfo>Reference>DigestValue"`
	SignatureValue         string             `xml:"SignatureValue"`
	KeyName                string             `xml:"KeyInfo>KeyName,omitempty"`
	X509Certificate        *SignatureX509Data `xml:"KeyInfo>X509Data,omitempty"`
}

// SignatureX509Data represents the <X509Data> element of <Signature>
type SignatureX509Data struct {
	X509Certificate string `xml:"X509Certificate,omitempty"`
}
