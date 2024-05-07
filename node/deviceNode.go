package node

type DeviceNode struct {
	name      string   // 设备名称
	vulnCount int      // 设备存在的漏洞总数
	vulnList  []string // 设备存在的漏洞列表
}

func (n DeviceNode) Hash() string {
	return n.name
}

func (n *DeviceNode) SetName(name string) {
	n.name = name
}

func (n DeviceNode) GetName() string {
	return n.name
}

func (n *DeviceNode) AddManyVuln(vulns []string) {
	n.vulnList = append(n.vulnList, vulns...)
	n.vulnCount = len(n.vulnList)
}

func (n DeviceNode) GetVulnList() []string {
	return n.vulnList
}

func (n DeviceNode) GetVulnCount() int {
	return n.vulnCount
}
