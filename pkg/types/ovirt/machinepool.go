package ovirt

// MachinePool stores the configuration for a machine pool installed
// on ovirt.
type MachinePool struct {
	// InstanceTypeID defines the VM instance type and overrides
	// the hardware parameters of the created VM, including cpu and memory.
	// If InstanceTypeID is passed, all memory and cpu variables will be ignored.
	InstanceTypeID string `json:"instanceTypeID,omitempty"`

	// CPU defines the VM CPU.
	CPU *CPU `json:"cpu,omitempty"`

	// MemoryMB is the size of a VM's memory in MiBs.
	MemoryMB int32 `json:"memoryMB,omitempty"`

	// OSDisk is the the root disk of the node.
	OSDisk *Disk `json:"osDisk,omitempty"`

	// VMType defines the workload type the instance will
	// be used for and this effects the instance parameters.
	// One of "desktop, server, high_performance"
	VMType string `json:"type,omitempty"`
}

// CPU defines the VM cpu, made of (Sockets * Cores * Threads)
type CPU struct {
	// Sockets is the number of sockets for a VM.
	// Total CPUs is (Sockets * Cores * Threads)
	Sockets int32 `json:"sockets"`

	// Cores is the number of cores per socket.
	// Total CPUs is (Sockets * Cores * Threads)
	Cores int32 `json:"cores"`
}

// Disk defines a VM disk
type Disk struct {
	// SizeGB size of the bootable disk in GiB.
	SizeGB int64 `json:"sizeGB"`
}

// Set sets the values from `required` to `p`.
func (p *MachinePool) Set(required *MachinePool) {
	if required == nil || p == nil {
		return
	}

	if required.InstanceTypeID != "" {
		p.InstanceTypeID = required.InstanceTypeID
	}

	if required.VMType != "" {
		p.VMType = required.VMType
	}

	if required.CPU != nil {
		p.CPU = required.CPU
	}

	if required.MemoryMB != 0 {
		p.MemoryMB = required.MemoryMB
	}

	if required.OSDisk != nil {
		p.OSDisk = required.OSDisk
	}
}
