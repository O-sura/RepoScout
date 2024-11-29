package internal

import "encoding/json"

// Filter out the repositories based on the best matching criterias
/*
Find suitable repos by looking @
 - most recent commit(Date)
 - No of open issues/PRs
 - No of commits
 - No of contributors
 - Age(When it is created)
 - No of stars
 - Contribution chart
 - Language
 - Tags(hacktoberfest)
 - Explore Issues: Using labels like 'good first issue', 'beginner-friendly', or 'help wanted', etc.(Have to make sure they are not too old)
*/
func FilterRepos(repoData json.RawMessage) {}
