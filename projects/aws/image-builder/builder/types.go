package builder

const (
	Ubuntu    string = "ubuntu"
	VSphere   string = "vsphere"
	Baremetal string = "baremetal"
	NutanixAHV       string = "nutanixahv"
)

type BuildOptions struct {
	Os              string
	Hypervisor      string
	VsphereConfig   string
	NutanixAHVConfig   string
	ReleaseChannel  string
	artifactsBucket string
	Force           bool
}
