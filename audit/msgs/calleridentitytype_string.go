// Code generated by "stringer -type=CallerIdentityType"; DO NOT EDIT.

package msgs

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownCallerIdentityType-0]
	_ = x[CIOther-1]
	_ = x[UPN-2]
	_ = x[PUID-3]
	_ = x[ObjectID-4]
	_ = x[Certificate-5]
	_ = x[Claim-6]
	_ = x[Username-7]
	_ = x[AccessKeyName-8]
	_ = x[SubscriptionID-9]
	_ = x[ApplicationID-10]
	_ = x[TenantID-11]
	_ = x[TokenID-12]
	_ = x[SasURL-13]
	_ = x[System-14]
}

const _CallerIdentityType_name = "UnknownCallerIdentityTypeCIOtherUPNPUIDObjectIDCertificateClaimUsernameAccessKeyNameSubscriptionIDApplicationIDTenantIDTokenIDSasURLSystem"

var _CallerIdentityType_index = [...]uint8{0, 25, 32, 35, 39, 47, 58, 63, 71, 84, 98, 111, 119, 126, 132, 138}

func (i CallerIdentityType) String() string {
	if i >= CallerIdentityType(len(_CallerIdentityType_index)-1) {
		return "CallerIdentityType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CallerIdentityType_name[_CallerIdentityType_index[i]:_CallerIdentityType_index[i+1]]
}
