package enum

var CODE = struct {
	SUCCESS string

	CUSTOMER_ERROR string

	SYSTEM_ERROR string

	THIRD_PARTY_ERROR string
}{
	SUCCESS: "00000",

	CUSTOMER_ERROR: "C0001",

	SYSTEM_ERROR: "S0001",

	THIRD_PARTY_ERROR: "T0001",
}
