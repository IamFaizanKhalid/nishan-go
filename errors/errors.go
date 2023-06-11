package errors

type Error int

const (
	Nil Error = 100

	FingerprintMismatch     Error = 111
	CitizenNumInvalid       Error = 112
	ServiceDown             Error = 113
	CitizenNumNotVerified   Error = 114
	FingerprintsUnavailable Error = 115
	FingerprintFormat       Error = 116
	FingerIndexUnavailable  Error = 117

	LongitudeInvalid Error = 119
	LatitudeInvalid  Error = 120
	ImeiInvalid      Error = 121
	OutsidePakistan  Error = 122
	MobileNumInvalid Error = 123

	FingerIndexInvalid      Error = 151
	FingerprintEmpty        Error = 152
	CitizenNumEmpty         Error = 153
	LimitReached            Error = 155 // for test environment only
	CitizenNumNotRegistered Error = 156 // for test environment only

	BalanceInsufficient Error = 171

	ApiKeyMissing Error = 181
	ApiKeyInvalid Error = 182

	TokenMissing        Error = 183
	TokenInvalid        Error = 184
	TokenExpired        Error = 185
	UnauthorizedService Error = 186

	SendingRequest Error = 500
)
