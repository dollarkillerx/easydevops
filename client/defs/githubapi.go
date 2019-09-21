/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:08 2019-09-16
 */
package defs

type GithubAPI struct {
	Branch     string `json:"ref" form:"ref"`
	Repository struct {
		FullName string `json:"full_name" form:"full_name"`
	}
	Commits []struct{
		Message string `json:"message" form:"message"`
	}
}

// "ref": "refs/heads/master",
// "ref": "refs/heads/test",
