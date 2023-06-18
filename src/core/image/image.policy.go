package image

import (
	"github.com/containers/image/v5/signature"
)

func GetImagePolicy() (*signature.PolicyContext, error) {
	var policy *signature.Policy
	var err error
	policy = &signature.Policy{Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}}
	if err != nil {
		return nil, err
	}
	return signature.NewPolicyContext(policy)
}
