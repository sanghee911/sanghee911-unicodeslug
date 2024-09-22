package unicodeslug

import "regexp"

func Slugify(input string) string {
	// replace all special characters including spaces with a hyphen
	reSpecialCharacters := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	slug := reSpecialCharacters.ReplaceAllString(input, "-")

	// replace a leading hyphen and an ending hyphen
	reHyphen := regexp.MustCompile(`^-+|-+$`)
	slug = reHyphen.ReplaceAllString(slug, "")
	return slug
}

func ValidateSlug(slug string) bool {
	re := regexp.MustCompile(`[!@#$%^&*()/\\\s]|^-|-$`)
	// Check if the slug contains any special characters excluding middle hyphens
	match := re.MatchString(slug)
	return !match
}
