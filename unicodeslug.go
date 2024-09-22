package unicodeslug

import "regexp"

func Slugify(input string) string {
	reSpecialCharacters := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	reHyphen := regexp.MustCompile(`^-+|-+$`)

	// replace all matches with a hyphen
	slug := reSpecialCharacters.ReplaceAllString(input, "-")
	// replace a leading hyphen and an ending hyphen
	slug = reHyphen.ReplaceAllString(slug, "")
	return slug
}
