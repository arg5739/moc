// Copyright (c) Microsoft Corporation.
// Licensed under the Apache v2.0 license.

package wssdcommon

import (
	pb "github.com/microsoft/moc/rpc/common"
)

type GpuAssignType = int32

const (
	GpuAssignTypeNone  GpuAssignType = 0
	GpuAssignTypeDDA   GpuAssignType = 1
	GpuAssignTypeGpuPv GpuAssignType = 2
)

type VmSize struct {
	CpuCount      int
	GpuCount      int
	GpuName       string
	MemoryMB      int
	GpuAssignMode GpuAssignType
}

// innerMap is captured in the closure returned below
var VirtualMachineSize_value = map[pb.VirtualMachineSizeType]VmSize{
	pb.VirtualMachineSizeType_Default: {
		CpuCount: 4,
		MemoryMB: 4096,
	},
	pb.VirtualMachineSizeType_Standard_A2_v2: {
		CpuCount: 2,
		MemoryMB: 4096,
	},
	pb.VirtualMachineSizeType_Standard_A4_v2: {
		CpuCount: 4,
		MemoryMB: 8192,
	},
	/// Standard DXs_v3
	pb.VirtualMachineSizeType_Standard_D2s_v3: {
		CpuCount: 2,
		MemoryMB: 8192,
	},
	pb.VirtualMachineSizeType_Standard_D4s_v3: {
		CpuCount: 4,
		MemoryMB: 16384,
	},
	pb.VirtualMachineSizeType_Standard_D8s_v3: {
		CpuCount: 8,
		MemoryMB: 32768,
	},
	pb.VirtualMachineSizeType_Standard_D16s_v3: {
		CpuCount: 16,
		MemoryMB: 65536,
	},
	pb.VirtualMachineSizeType_Standard_D32s_v3: {
		CpuCount: 32,
		MemoryMB: 131072,
	},
	/// Standard DSX_v2
	pb.VirtualMachineSizeType_Standard_DS2_v2: {
		CpuCount: 2,
		MemoryMB: 7168,
	},
	pb.VirtualMachineSizeType_Standard_DS3_v2: {
		CpuCount: 2,
		MemoryMB: 14336,
	},
	pb.VirtualMachineSizeType_Standard_DS4_v2: {
		CpuCount: 8,
		MemoryMB: 28672,
	},
	pb.VirtualMachineSizeType_Standard_DS5_v2: {
		CpuCount: 16,
		MemoryMB: 57344,
	},
	/// Standard DSX_v2 (memory optimized)
	pb.VirtualMachineSizeType_Standard_DS13_v2: {
		CpuCount: 8,
		MemoryMB: 57344,
	},
	/// Custom sizes for MSK8S
	pb.VirtualMachineSizeType_Standard_K8S_v1: {
		CpuCount: 4,
		MemoryMB: 2048,
	},
	pb.VirtualMachineSizeType_Standard_K8S2_v1: {
		CpuCount: 2,
		MemoryMB: 2048,
	},
	pb.VirtualMachineSizeType_Standard_K8S3_v1: {
		CpuCount: 4,
		MemoryMB: 6144,
	},
	pb.VirtualMachineSizeType_Standard_K8S4_v1: {
		CpuCount: 4,
		MemoryMB: 4096,
	},
	pb.VirtualMachineSizeType_Standard_K8S5_v1: {
		CpuCount: 2,
		MemoryMB: 1024,
	},
	// Sizes with GPUs
	pb.VirtualMachineSizeType_Standard_NK6: {
		CpuCount:      6,
		GpuCount:      1,
		GpuName:       "NVIDIA Tesla T4",
		MemoryMB:      12288,
		GpuAssignMode: GpuAssignTypeDDA,
	},
	pb.VirtualMachineSizeType_Standard_NK12: {
		CpuCount:      12,
		GpuCount:      2,
		GpuName:       "NVIDIA Tesla T4",
		MemoryMB:      24576,
		GpuAssignMode: GpuAssignTypeDDA,
	},
	pb.VirtualMachineSizeType_Standard_NV6: {
		CpuCount:      6,
		GpuCount:      1,
		GpuName:       "NVIDIA Tesla M60",
		MemoryMB:      57344,
		GpuAssignMode: GpuAssignTypeDDA,
	},
	pb.VirtualMachineSizeType_Standard_NV12: {
		CpuCount:      12,
		GpuCount:      2,
		GpuName:       "NVIDIA Tesla M60",
		MemoryMB:      131072,
		GpuAssignMode: GpuAssignTypeDDA,
	},
}
