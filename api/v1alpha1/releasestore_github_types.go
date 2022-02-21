package v1alpha1

type GithubProvider struct {
    Repository string `json:"repository"`
    // +optional
    TagRegex   string `json:"tagRegex"`
}
