package errors

type ErrorCode int

const (
	Nil ErrorCode = 100

	FingerprintMismatch     ErrorCode = 111
	CitizenNumInvalid       ErrorCode = 112
	ServiceDown             ErrorCode = 113
	CitizenNumNotVerified   ErrorCode = 114
	FingerprintsUnavailable ErrorCode = 115
	FingerprintFormat       ErrorCode = 116
	FingerIndexUnavailable  ErrorCode = 117

	LongitudeInvalid ErrorCode = 119
	LatitudeInvalid  ErrorCode = 120
	ImeiInvalid      ErrorCode = 121
	OutsidePakistan  ErrorCode = 122
	MobileNumInvalid ErrorCode = 123

	FingerIndexInvalid ErrorCode = 151
	FingerprintEmpty   ErrorCode = 152
	CitizenNumEmpty    ErrorCode = 153

	LimitReached            ErrorCode = 155 // for test environment only
	CitizenNumNotRegistered ErrorCode = 156 // for test environment only

	BalanceInsufficient ErrorCode = 171

	ApiKeyMissing ErrorCode = 181
	ApiKeyInvalid ErrorCode = 182

	TokenMissing        ErrorCode = 183
	TokenInvalid        ErrorCode = 184
	TokenExpired        ErrorCode = 185
	UnauthorizedService ErrorCode = 186

	SendingRequest ErrorCode = 500
)
