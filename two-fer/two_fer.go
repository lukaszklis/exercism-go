// Package twofer should have a package comment that summarizes what it's about.
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
