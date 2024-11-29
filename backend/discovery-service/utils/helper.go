package utils

// HasTag checks if a repository contains any of the target tags.
func HasTag(repoTags []string, targetTags []string) bool {
	for _, tag := range repoTags {
		for _, target := range targetTags {
			if tag == target {
				return true
			}
		}
	}
	return false
}
