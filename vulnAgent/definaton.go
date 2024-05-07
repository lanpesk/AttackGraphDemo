package vulnagent

// the defination is refer from the NVD

const USERPRIVILEGE = 0
const ROOTPRIVILEGE = 1
const NONEPRIVILEGE = 2
const _NVD_REST_API = "https://services.nvd.nist.gov/rest/json/cves/2.0"
const ERROR = -1

/*
deprecated

// this is for the cvss version
const VERSION2 = 0
const VERSION3 = 1
const VERSION31 = 2

// this is for accessVector
const AV_NETWORK = 3
const AV_ADJACENT_NETWORK = 4
const AV_LOCAL = 5
const AV_PHYSICAL = 6

// attention: in cvss3.1 it doesnt use MEDIUM level
const AC_LOW = 7
const AC_MEDIUM = 8
const AC_HIGH = 9

// this is only in cvss3
const PR_LOW = 10
const PR_HIGH = 11
const UI_NONE = 12
const UI_REQUIRED = 13
const S_UNCHANGED = 14
const S_CHANGED = 15

// this is for authentication
const AU_NONE = 16
const AU_SINGLE = 17
const AU_MUL = 18

const CVSS2_C_NONE = 19
const CVSS2_C_PARTIAL = 20
const CVSS2_C_COMPLETE = 21

const CVSS2_I_NONE = 22
const CVSS2_I_PARTIAL = 23
const CVSS2_I_COMPLETE = 24

const CVSS2_A_NONE = 25
const CVSS2_A_PARTIAL = 26
const CVSS2_A_COMPLETE = 27

// below is cvss3.1
const CVSS31_C_NONE = 28
const CVSS31_C_LOW = 29
const CVSS31_C_HIGH = 30

const CVSS31_I_NONE = 31
const CVSS31_I_LOW = 32
const CVSS31_I_HIGH = 33

const CVSS31_A_NONE = 34
const CVSS31_A_LOW = 35
const CVSS31_A_HIGH = 36

const SERERITY_NONE = 37
const SERERITY_LOW = 38
const SERERITY_MEDIUM = 39
const SERERITY_HIGH = 40
const SERERITY_CRITICAL = 41

*/
