package ident

func IsValid(ident string) bool {
	if len(ident) == 0 {
		return false
	}

	for i, r := range ident {
		switch {
		case r == '_':
			continue
		case 'a' <= r && r <= 'z':
			continue
		case 'A' <= r && r <= 'Z':
			continue
		case '0' <= r && r <= '9' && i > 0:
			continue
		default:
			return false
		}
	}
	return true
}

func IsAnonymous(ident string) bool {
	return ident == "_"
}

func IsPublic(ident string) bool {
	if !IsValid(ident) {
		return false
	}

	r := ident[0]
	return 'A' <= r && r <= 'Z'
}
