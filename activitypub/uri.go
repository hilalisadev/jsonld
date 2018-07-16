package activitypub

type (
	// IRI is a Internationalized Resource Identifiers (IRIs) RFC3987
	IRI URI

	// URI is a Uniform Resource Identifier (URI) RFC3986
	URI string
)

// String returns the String value of the URI object
func (u URI) String() string {
	return string(u)
}

// String returns the String value of the IRI object
func (i IRI) String() string {
	return string(i)
}

// GetLink returns a copy of itself
func (u URI) GetLink() URI {
	return u
}

// GetLink returns a URI type coercion of the IRI object
func (i IRI) GetLink() URI {
	return URI(i)
}
