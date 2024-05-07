package vulnagent

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

type CVE struct {
	Name         string         `json:"cve-id"` // 漏洞名称
	Cvss2Metric  *CVSSv2Metric  `json:"cvss2metric"`
	Cvss31Metric *CVSSv31Metric `json:"cvss31metric"`
}

type CVSSv2Metric struct {
	Source                  string     `json:"source"`
	Type                    string     `json:"type"`
	CvssData                CVSSv2Data `json:"cvssData"`
	BaseServerity           string     `json:"baseSeverity"`
	ExploitablilityScore    float32    `json:"exploitabilityScore"`
	ImpactScore             float32    `json:"impactScore"`
	AcInsufInfo             bool       `json:"acInsufInfo"`
	ObtainAllPrivilege      bool       `json:"obtainAllPrivilege"`
	ObtainUserPrivilege     bool       `json:"obtainUserPrivilege"`
	ObtainOtherPrivilege    bool       `json:"obtainOtherPrivilege"`
	UserInteractionRequired bool       `json:"userInteractionRequired"`
}

type CVSSv2Data struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	AccessVector          string  `json:"accessVector"`
	AccessComplexity      string  `json:"accessComplexity"`
	Authentication        string  `json:"authentication"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntergrityImpact      string  `json:"integrityImpact"`
	AvailablilityImpact   string  `json:"availabilityImpact"`
	BaseScore             float32 `json:"baseScore"`
}

type CVSSv31Metric struct {
	Source               string      `json:"source"`
	Type                 string      `json:"type"`
	CvssData             CVSSv31Data `json:"cvssData"`
	ExploitablilityScore float32     `json:"exploitabilityScore"`
	ImpactScore          float32     `json:"impactScore"`
}

type CVSSv31Data struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	AttackVector          string  `json:"attackVector"`
	AttackComplexity      string  `json:"attackComplexity"`
	PrivilegesRequired    string  `json:"privilegesRequired"`
	UserInteraction       string  `json:"userInteraction"`
	Scope                 string  `json:"scope"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntergrityImpact      string  `json:"integrityImpact"`
	AvailablilityImpact   string  `json:"availabilityImpact"`
	BaseScore             float32 `json:"baseScore"`
	BaseServerity         string  `json:"baseSeverity"`
}

// type ListCVE struct {
// 	DataBase string         // 漏洞来源,有NVD, CVE等
// 	list     map[string]CVE // 漏洞字典
// }

type VulnAgentError struct {
	info string
}

func (e *VulnAgentError) Error() string {
	return e.info
}

// using for get json text from rest api.
func wget(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *CVE) QueryVulnInfo() error {
	if c.Name == "" {
		return &VulnAgentError{"The CVE ID is not set"}
	}

	data, err := wget(_NVD_REST_API + "?cveId=" + c.Name)
	if err != nil {
		return err
	}

	r := gjson.GetBytes(data, "vulnerabilities.0.cve.metrics.cvssMetricV2.0")
	if r.Exists() {
		t := new(CVSSv2Metric)
		err = json.Unmarshal([]byte(r.Raw), t)
		if err != nil {
			return err
		}
		c.Cvss2Metric = t

	} else {
		c.Cvss2Metric = nil
	}

	r = gjson.GetBytes(data, "vulnerabilities.0.cve.metrics.cvssMetricV31.0")
	if r.Exists() {
		t := new(CVSSv31Metric)
		err = json.Unmarshal([]byte(r.Raw), t)
		if err != nil {
			return err
		}
		c.Cvss31Metric = t

	} else {
		c.Cvss31Metric = nil
	}

	return nil
}
