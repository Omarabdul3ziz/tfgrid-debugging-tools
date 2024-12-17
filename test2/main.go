package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/peer"
	"github.com/threefoldtech/zos/client"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

const (
	Substrate = "wss://tfchain.dev.grid.tf:443"
	NodeID    = 259
	Twin      = 29
)

// func network() gridtypes.Workload {
// 	wgKey := "GDU+cjKrHNJS9fodzjFDzNFl5su3kJXTZ3ipPgUjOUE="

// 	return gridtypes.Workload{
// 		Version:     Version,
// 		Type:        zos.NetworkType,
// 		Description: "test network",
// 		Name:        "network",
// 		Data: gridtypes.MustMarshal(zos.Network{
// 			NetworkIPRange: gridtypes.MustParseIPNet("10.1.0.0/16"),
// 			Subnet:         gridtypes.MustParseIPNet("10.1.1.0/24"),
// 			WGPrivateKey:   wgKey,
// 			WGListenPort:   3011,
// 			Peers: []zos.Peer{
// 				{
// 					Subnet:      gridtypes.MustParseIPNet("10.1.2.0/24"),
// 					WGPublicKey: "4KTvZS2KPWYfMr+GbiUUly0ANVg8jBC7xP9Bl79Z8zM=",

// 					AllowedIPs: []gridtypes.IPNet{
// 						gridtypes.MustParseIPNet("10.1.2.0/24"),
// 						gridtypes.MustParseIPNet("100.64.1.2/32"),
// 					},
// 				},
// 				// {
// 				// 	Subnet:      gridtypes.MustParseIPNet("10.1.3.0/24"),
// 				// 	WGPublicKey: "NnrqoeJUUWIgW/oJzl3MBu+KeVv8pYhGrlLmvdFcyUw=",

// 				// 	AllowedIPs: []gridtypes.IPNet{
// 				// 		gridtypes.MustParseIPNet("10.1.3.0/24"),
// 				// 		gridtypes.MustParseIPNet("100.64.1.3/32"),
// 				// 	},
// 				// },
// 			},
// 		}),
// 	}
// }

func networkLight() gridtypes.Workload {
	// b, _ := hex.DecodeString("9751c596c7c951aedad1a5f78f18b59515064adf660e0d55abead65e6fbbd627")
	return gridtypes.Workload{
		Version:     Version,
		Type:        zos.NetworkType,
		Description: "testnetwork",
		Name:        "netabcdfg",
		Data: gridtypes.MustMarshal(zos.Network{
			Subnet: gridtypes.MustParseIPNet("10.1.1.0/24"),
			// Mycelium: zos.Mycelium{
			// 	Key: b,
			// },
			NetworkIPRange: gridtypes.MustParseIPNet("10.1.0.0/16"),
		}),
	}
}

func zdbLight(size gridtypes.Unit) gridtypes.Workload {
	return gridtypes.Workload{
		Version: Version,
		Type:    zos.ZDBType,
		Name:    "zdb1211",
		Data: gridtypes.MustMarshal(zos.ZDB{
			Size:     size,
			Mode:     zos.ZDBModeUser,
			Password: "password",
		},
		),
	}
}

func disk(name string, size gridtypes.Unit) gridtypes.Workload {
	return gridtypes.Workload{
		Name:        gridtypes.Name(name),
		Version:     Version,
		Type:        zos.ZMountType,
		Description: "test",
		Data: gridtypes.MustMarshal(zos.ZMount{
			Size: size + 1,
		}),
	}
}

func network(name gridtypes.Name) gridtypes.Workload {
	wgKey := "GDU+cjKrHNJS9fodzjFDzNFl5su3kJXTZ3ipPgUjOUE="

	return gridtypes.Workload{
		Version:     Version,
		Type:        zos.NetworkType,
		Description: "test network",
		Name:        name,
		Data: gridtypes.MustMarshal(zos.Network{
			NetworkIPRange: gridtypes.MustParseIPNet("10.1.0.0/16"),
			Subnet:         gridtypes.MustParseIPNet("10.1.1.0/24"),
			WGPrivateKey:   wgKey,
			WGListenPort:   3011,
			Peers: []zos.Peer{
				{
					Subnet:      gridtypes.MustParseIPNet("10.1.2.0/24"),
					WGPublicKey: "4KTvZS2KPWYfMr+GbiUUly0ANVg8jBC7xP9Bl79Z8zM=",

					AllowedIPs: []gridtypes.IPNet{
						gridtypes.MustParseIPNet("10.1.2.0/24"),
						gridtypes.MustParseIPNet("100.64.1.2/32"),
					},
				},
				// {
				//   Subnet:      gridtypes.MustParseIPNet("10.1.3.0/24"),
				//   WGPublicKey: "NnrqoeJUUWIgW/oJzl3MBu+KeVv8pYhGrlLmvdFcyUw=",

				//   AllowedIPs: []gridtypes.IPNet{
				//     gridtypes.MustParseIPNet("10.1.3.0/24"),
				//     gridtypes.MustParseIPNet("100.64.1.3/32"),
				//   },
				// },
			},
		}),
	}
}

func zdb(name string, size gridtypes.Unit) gridtypes.Workload {
	return gridtypes.Workload{
		Version:     Version,
		Name:        gridtypes.Name(name),
		Type:        zos.ZDBType,
		Description: "test zdb 2",
		Data: gridtypes.MustMarshal(zos.ZDB{
			Size:     size,
			Mode:     zos.ZDBModeUser,
			Password: "password",
		}),
	}
}

const Version = 0

func publicip(name string) gridtypes.Workload {
	return gridtypes.Workload{
		Version: Version,
		Name:    gridtypes.Name(name),
		Type:    zos.PublicIPType,
		Data:    gridtypes.MustMarshal(zos.PublicIP{V4: true, V6: false}),
	}
}

//	func container(mounts map[string]string) gridtypes.Workload {
//		return gridtypes.Workload{
//			Version: Version,
//			Name:    "container",
//			Type:    zos.ZMachineType,
//			Data: gridtypes.MustMarshal(zos.ZMachine{
//				// FList: "http://185.69.167.145/flists/base-host.fl",
//				FList: "https://hub.grid.tf/tf-official-apps/threefoldtech-ubuntu-22.04.flist",
//				// FList: "https://hub.grid.tf/azmy.3bot/ubuntu-jammy.flist",
//				// FList: "https://hub.grid.tf/aelawady.3bot/alpine-ssh.flist",
//				// FList: "https://hub.grid.tf/petep.3bot/ubuntu-focal.flist",
//				Network: zos.MachineNetwork{
//					Interfaces: []zos.MachineInterface{
//						{
//							Network: "network",
//							IP:      net.ParseIP("10.1.1.0"),
//						},
//					},
//					Planetary: true,
//				},
//				Size: 5 * gridtypes.Gigabyte,
//				ComputeCapacity: zos.MachineCapacity{
//					CPU:    1,
//					Memory: 512 * gridtypes.Megabyte,
//				},
//				Entrypoint: "/sbin/zinit init",
//				Mounts: func() []zos.MachineMount {
//					var mnt []zos.MachineMount
//					for k, v := range mounts {
//						mnt = append(mnt, zos.MachineMount{
//							Name:       gridtypes.Name(k),
//							Mountpoint: v,
//						})
//					}
//					return mnt
//				}(),
//				Env: map[string]string{
//					"SSH_KEY": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDcGrS1RT36rHAGLK3/4FMazGXjIYgWVnZ4bCvxxg8KosEEbs/DeUKT2T2LYV91jUq3yibTWwK0nc6O+K5kdShV4qsQlPmIbdur6x2zWHPeaGXqejbbACEJcQMCj8szSbG8aKwH8Nbi8BNytgzJ20Ysaaj2QpjObCZ4Ncp+89pFahzDEIJx2HjXe6njbp6eCduoA+IE2H9vgwbIDVMQz6y/TzjdQjgbMOJRTlP+CzfbDBb6Ux+ed8F184bMPwkFrpHs9MSfQVbqfIz8wuq/wjewcnb3wK9dmIot6CxV2f2xuOZHgNQmVGratK8TyBnOd5x4oZKLIh3qM9Bi7r81xCkXyxAZbWYu3gGdvo3h85zeCPGK8OEPdYWMmIAIiANE42xPmY9HslPz8PAYq6v0WwdkBlDWrG3DD3GX6qTt9lbSHEgpUP2UOnqGL4O1+g5Rm9x16HWefZWMjJsP6OV70PnMjo9MPnH+yrBkXISw4CGEEXryTvupfaO5sL01mn+UOyE= abdulrahman@AElawady-PC",
//				},
//			}),
//		}
//	}

func containerLight() gridtypes.Workload {
	b, _ := hex.DecodeString("b60f2b7ec39c")
	return gridtypes.Workload{
		Version: Version,
		Name:    "container3",
		Type:    zos.ZMachineType,
		Data: gridtypes.MustMarshal(zos.ZMachine{
			// FList: "http://185.69.167.145/flists/base-host.fl",
			FList: "https://hub.grid.tf/tf-official-vms/ubuntu-24.04-full.flist",
			// FList: "https://hub.grid.tf/azmy.3bot/ubuntu-jammy.flist",
			// FList: "https://hub.grid.tf/aelawady.3bot/alpine-ssh.flist",
			// FList: "https://hub.grid.tf/petep.3bot/ubuntu-focal.flist",
			Network: zos.MachineNetwork{
				Interfaces: []zos.MachineInterface{
					{
						Network: "netabc",
						IP:      net.ParseIP("10.1.1.2"),
					},
				},
				Mycelium: &zos.MyceliumIP{
					Network: "netabc",
					Seed:    b,
				},
			},
			Size: 1 * gridtypes.Gigabyte,
			ComputeCapacity: zos.MachineCapacity{
				CPU:    1,
				Memory: 512 * gridtypes.Megabyte,
			},
			Entrypoint: "/sbin/zinit init",
			Env: map[string]string{
				"SSH_KEY": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDJ1t4Ug8EfykmJwAbYudyYYN/f7dZaVg3KGD2Pz0bd9pajAAASWYrss3h2ctCZWluM6KAt289RMNzxlNUkOMJ9WhCIxqDAwtg05h/J27qlaGCPP8BCEITwqNKsLwzmMZY1UFc+sSUyjd35d3kjtv+rzo2meaReZnUFNPisvxGoygftAE6unqNa7TKonVDS1YXzbpT8XdtCV1Y6ACx+3a82mFR07zgmY4BVOixNBy2Lzpq9KiZTz91Bmjg8dy4xUyWLiTmnye51hEBgUzPprjffZByYSb2Ag9hpNE1AdCGCli/0TbEwFn9iEroh/xmtvZRpux+L0OmO93z5Sz+RLiYXKiYVV5R5XYP8y5eYi48RY2qr82sUl5+WnKhI8nhzayO9yjPEp3aTvR1FdDDj5ocB7qKi47R8FXIuwzZf+kJ7ZYmMSG7N21zDIJrz6JGy9KMi7nX1sqy7NSqX3juAasIjx0IJsE8zv9qokZ83hgcDmTJjnI+YXimelhcHn4M52hU= omar@jarvis",
			},
		}),
	}
}

// func containerLight2() gridtypes.Workload {
// 	b, _ := hex.DecodeString("b62f1b7ec39c")
// 	return gridtypes.Workload{
// 		Version: Version,
// 		Name:    "container2",
// 		Type:    zos.ZMachineLightType,
// 		Data: gridtypes.MustMarshal(zos.ZMachineLight{
// 			// FList: "http://185.69.167.145/flists/base-host.fl",
// 			FList: "https://hub.grid.tf/tf-official-apps/threefoldtech-ubuntu-22.04.flist",
// 			// FList: "https://hub.grid.tf/azmy.3bot/ubuntu-jammy.flist",
// 			// FList: "https://hub.grid.tf/aelawady.3bot/alpine-ssh.flist",
// 			// FList: "https://hub.grid.tf/petep.3bot/ubuntu-focal.flist",
// 			Network: zos.MachineNetworkLight{
// 				Interfaces: []zos.MachineInterface{
// 					{
// 						Network: "netabc",
// 						IP:      net.ParseIP("10.1.1.3"),
// 					},
// 				},
// 				Mycelium: &zos.MyceliumIP{
// 					Network: "netabc",
// 					Seed:    b,
// 				},
// 			},
// 			Size: 1 * gridtypes.Gigabyte,
// 			ComputeCapacity: zos.MachineCapacity{
// 				CPU:    1,
// 				Memory: 512 * gridtypes.Megabyte,
// 			},
// 			Entrypoint: "/sbin/zinit init",
// 			Env: map[string]string{
// 				"SSH_KEY": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDJ1t4Ug8EfykmJwAbYudyYYN/f7dZaVg3KGD2Pz0bd9pajAAASWYrss3h2ctCZWluM6KAt289RMNzxlNUkOMJ9WhCIxqDAwtg05h/J27qlaGCPP8BCEITwqNKsLwzmMZY1UFc+sSUyjd35d3kjtv+rzo2meaReZnUFNPisvxGoygftAE6unqNa7TKonVDS1YXzbpT8XdtCV1Y6ACx+3a82mFR07zgmY4BVOixNBy2Lzpq9KiZTz91Bmjg8dy4xUyWLiTmnye51hEBgUzPprjffZByYSb2Ag9hpNE1AdCGCli/0TbEwFn9iEroh/xmtvZRpux+L0OmO93z5Sz+RLiYXKiYVV5R5XYP8y5eYi48RY2qr82sUl5+WnKhI8nhzayO9yjPEp3aTvR1FdDDj5ocB7qKi47R8FXIuwzZf+kJ7ZYmMSG7N21zDIJrz6JGy9KMi7nX1sqy7NSqX3juAasIjx0IJsE8zv9qokZ83hgcDmTJjnI+YXimelhcHn4M52hU= omar@jarvis",
// 			},
// 		}),
// 	}
// }

// func vm(mounts map[string]string, ip string) gridtypes.Workload {
// 	return gridtypes.Workload{
// 		Version: Version,
// 		Name:    "vm",
// 		Type:    zos.ZMachineType,
// 		Data: gridtypes.MustMarshal(zos.ZMachine{
// 			//FList: "https://hub.grid.tf/azmy.3bot/ubuntujammy.flist",
// 			//FList: "https://hub.grid.tf/azmy.3bot/ubuntu-22.04.flist",
// 			FList: "https://hub.grid.tf/tf-official-vms/ubuntu-22.04-lts.flist",
// 			// FList: "https://hub.grid.tf/tf-official-vms/archlinux.flist",
// 			Network: zos.MachineNetwork{
// 				Interfaces: []zos.MachineInterface{
// 					{
// 						Network: "network",
// 						IP:      net.ParseIP("10.1.1.3"),
// 					},
// 				},
// 				Planetary: true,
// 				PublicIP:  gridtypes.Name(ip),
// 			},
// 			// ComputeCapacity: zos.MachineCapacity{
// 			// 	CPU:    2,
// 			// 	Memory: 512 * gridtypes.Megabyte,
// 			// },
// 			Mounts: func() []zos.MachineMount {
// 				var mnt []zos.MachineMount
// 				for k, v := range mounts {
// 					mnt = append(mnt, zos.MachineMount{
// 						Name:       gridtypes.Name(k),
// 						Mountpoint: v,
// 					})
// 				}
// 				return mnt
// 			}(),
// 			Env: map[string]string{
// 				"SSH_KEY": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQJpNHC/AWNRAuDkqAZH6Y6pJsafTT9kslY486v0Bt2PfVK12mGdzLblzneqKdL/c75XUC4ujxR/7RXdGV+bicMoFJDFJeXCGyglzq0ep86TSMnw8/17uLWHunxzHs3xMSAXVZnCHeOB9EMvkInn0SS6Bp6SfkVDcx1kVFoY4+UXI4+/OQDkzbP6BB1QUcexeyqknAhFFaB6xCoMajRSgwoGTbEmc2dIc3jT5FJyW4WxEUhbI3cFd/LmVwVp5ttEVoW7sWUEHcG6CFg6NUVOkcpQc0X7YuBJ7oNgZFSKyUiumQO54ABtmzSOovUal0/GCNblv9nka8sfyod5DAYofbPGNrqKHnDkJRk9dQeB2xuRNK2Uiyz/iw/f13qc7WXdPeYHUhz3HSsn6EX0+wWK+0Sbk5kdVd+Hl8T3Ra7O1e7p5JuAUjcYtrBdw1KE3JxXjnpH33ORKj9Y/obyVcbjvIrMTf0JjGoG76DQFS+j5dRlfVcf0Ldb194PsqYCbwAUs=",
// 			},
// 			// GPU: []zos.GPU{
// 			// 	"0000:0e:00.0/1002/744c",
// 			// },
// 		}),
// 	}
// }

// func containerPublic(ip string) gridtypes.Workload {
// 	return gridtypes.Workload{
// 		Version: Version,
// 		Name:    "vm",
// 		Type:    zos.ZMachineType,
// 		Data: gridtypes.MustMarshal(zos.ZMachine{
// 			//FList: "https://hub.grid.tf/tf-official-apps/base:latest.flist",
// 			FList: "https://hub.grid.tf/azmy.3bot/forwarder.flist",
// 			Network: zos.MachineNetwork{
// 				Interfaces: []zos.MachineInterface{
// 					{
// 						Network: "network",
// 						IP:      net.ParseIP("10.1.1.3"),
// 					},
// 				},
// 				Planetary: true,
// 				PublicIP:  gridtypes.Name(ip),
// 			},
// 			ComputeCapacity: zos.MachineCapacity{
// 				CPU:    1,
// 				Memory: 512 * gridtypes.Megabyte,
// 			},
// 			//Entrypoint: "/sbin/zinit init",
// 			Env: map[string]string{
// 				"SSH_KEY": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQJpNHC/AWNRAuDkqAZH6Y6pJsafTT9kslY486v0Bt2PfVK12mGdzLblzneqKdL/c75XUC4ujxR/7RXdGV+bicMoFJDFJeXCGyglzq0ep86TSMnw8/17uLWHunxzHs3xMSAXVZnCHeOB9EMvkInn0SS6Bp6SfkVDcx1kVFoY4+UXI4+/OQDkzbP6BB1QUcexeyqknAhFFaB6xCoMajRSgwoGTbEmc2dIc3jT5FJyW4WxEUhbI3cFd/LmVwVp5ttEVoW7sWUEHcG6CFg6NUVOkcpQc0X7YuBJ7oNgZFSKyUiumQO54ABtmzSOovUal0/GCNblv9nka8sfyod5DAYofbPGNrqKHnDkJRk9dQeB2xuRNK2Uiyz/iw/f13qc7WXdPeYHUhz3HSsn6EX0+wWK+0Sbk5kdVd+Hl8T3Ra7O1e7p5JuAUjcYtrBdw1KE3JxXjnpH33ORKj9Y/obyVcbjvIrMTf0JjGoG76DQFS+j5dRlfVcf0Ldb194PsqYCbwAUs=",
// 				"TARGET":  "100.64.1.2",
// 			},
// 		}),
// 	}

// return gridtypes.Workload{
// 	Version: Version,
// 	Name:    "public",
// 	Type:    zos.ZMachineType,
// 	Data: gridtypes.MustMarshal(zos.ZMachine{
// 		FList: "https://hub.grid.tf/tf-official-apps/base:latest.flist",
// 		Network: zos.MachineNetwork{
// 			Interfaces: []zos.MachineInterface{
// 				{
// 					Network: "network",
// 					IP:      net.ParseIP("10.1.1.4"),
// 				},
// 			},
// 			Planetary: true,
// 			PublicIP:  "pub",
// 		},
// 		ComputeCapacity: zos.MachineCapacity{
// 			CPU:    1,
// 			Memory: 250 * gridtypes.Megabyte,
// 		},
// 		Entrypoint: "/sbin/zinit init",
// 		Env: map[string]string{
// 			"SSH_KEY": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQJpNHC/AWNRAuDkqAZH6Y6pJsafTT9kslY486v0Bt2PfVK12mGdzLblzneqKdL/c75XUC4ujxR/7RXdGV+bicMoFJDFJeXCGyglzq0ep86TSMnw8/17uLWHunxzHs3xMSAXVZnCHeOB9EMvkInn0SS6Bp6SfkVDcx1kVFoY4+UXI4+/OQDkzbP6BB1QUcexeyqknAhFFaB6xCoMajRSgwoGTbEmc2dIc3jT5FJyW4WxEUhbI3cFd/LmVwVp5ttEVoW7sWUEHcG6CFg6NUVOkcpQc0X7YuBJ7oNgZFSKyUiumQO54ABtmzSOovUal0/GCNblv9nka8sfyod5DAYofbPGNrqKHnDkJRk9dQeB2xuRNK2Uiyz/iw/f13qc7WXdPeYHUhz3HSsn6EX0+wWK+0Sbk5kdVd+Hl8T3Ra7O1e7p5JuAUjcYtrBdw1KE3JxXjnpH33ORKj9Y/obyVcbjvIrMTf0JjGoG76DQFS+j5dRlfVcf0Ldb194PsqYCbwAUs=",
// 		},
// 	}),
// }
// }

// func gateway(backend string, network *gridtypes.Name) gridtypes.Workload {
// 	return gridtypes.Workload{
// 		Version: Version,
// 		Name:    "gateway",
// 		Type:    zos.GatewayNameProxyType,
// 		Data: gridtypes.MustMarshal(zos.GatewayNameProxy{
// 			Name: "zuzu",
// 			GatewayBase: zos.GatewayBase{
// 				Backends: []zos.Backend{
// 					zos.Backend(backend),
// 				},
// 				Network: network,
// 			},
// 		}),
// 	}
// }

func logs() gridtypes.Workload {
	return gridtypes.Workload{
		Version: Version,
		Name:    "logs",
		Type:    zos.ZLogsType,
		Data: gridtypes.MustMarshal(zos.ZLogs{
			ZMachine: "vm",
			Output:   "ws://192.168.123.1:8080/logs/my-vm-logs",
		}),
	}
}

func vm(mounts map[string]string, ip string) gridtypes.Workload {
	return gridtypes.Workload{
		Version:     Version,
		Name:        "vm2",
		Description: "zmachine test",
		Type:        zos.ZMachineType,
		Data: gridtypes.MustMarshal(zos.ZMachine{
			// FList: "https://hub.grid.tf/azmy.3bot/ubuntujammy.flist",
			// FList: "https://hub.grid.tf/azmy.3bot/ubuntu-22.04.flist",
			FList: "https://hub.grid.tf/tf-official-vms/ubuntu-24.04-full.flist",
			// FList: "https://hub.grid.tf/tf-official-vms/archlinux.flist",
			Network: zos.MachineNetwork{
				Interfaces: []zos.MachineInterface{
					{
						Network: "network2",
						IP:      net.ParseIP("10.1.1.3"),
					},
				},
				Planetary: true,
				PublicIP:  gridtypes.Name(ip),
			},
			ComputeCapacity: zos.MachineCapacity{
				CPU:    1,
				Memory: 256 * gridtypes.Megabyte,
			},
			Mounts: func() []zos.MachineMount {
				var mnt []zos.MachineMount
				for k, v := range mounts {
					mnt = append(mnt, zos.MachineMount{
						Name:       gridtypes.Name(k),
						Mountpoint: v,
					})
				}
				return mnt
			}(),
			Env: map[string]string{
				"SSH_KEY": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDJ1t4Ug8EfykmJwAbYudyYYN/f7dZaVg3KGD2Pz0bd9pajAAASWYrss3h2ctCZWluM6KAt289RMNzxlNUkOMJ9WhCIxqDAwtg05h/J27qlaGCPP8BCEITwqNKsLwzmMZY1UFc+sSUyjd35d3kjtv+rzo2meaReZnUFNPisvxGoygftAE6unqNa7TKonVDS1YXzbpT8XdtCV1Y6ACx+3a82mFR07zgmY4BVOixNBy2Lzpq9KiZTz91Bmjg8dy4xUyWLiTmnye51hEBgUzPprjffZByYSb2Ag9hpNE1AdCGCli/0TbEwFn9iEroh/xmtvZRpux+L0OmO93z5Sz+RLiYXKiYVV5R5XYP8y5eYi48RY2qr82sUl5+WnKhI8nhzayO9yjPEp3aTvR1FdDDj5ocB7qKi47R8FXIuwzZf+kJ7ZYmMSG7N21zDIJrz6JGy9KMi7nX1sqy7NSqX3juAasIjx0IJsE8zv9qokZ83hgcDmTJjnI+YXimelhcHn4M52hU= omar@jarvis",
			},
		}),
	}
}

func main() {
	// const NodeID = "B7vj97BrXCp7yENj5zFN5ecZaqwSJEUmsJ8Naf8rzokR"
	// const DevNodeID = "3We9p7C3wihY3ZZSP5w4ujEBPVtm685fVW563MVFbnE8"

	Mnemonic := "mom picnic deliver again rug night rabbit music motion hole lion where"
	identity, err := substrate.NewIdentityFromSr25519Phrase(Mnemonic)
	if err != nil {
		panic(err)
	}

	// seed, err := bip39.EntropyFromMnemonic(Mnemonics)
	// if err != nil {
	// 	panic(err)
	// }
	// seed, err := hex.DecodeString(Seed)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("required:", ed25519.SeedSize, "given:", len(seed))
	// net := gridtypes.Name("network")
	dl := gridtypes.Deployment{
		Version:     Version,
		TwinID:      Twin, // LocalTwin,
		Description: "newaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Metadata:    "eoiwefo",
		// this contract id must match the one on substrate
		Workloads: []gridtypes.Workload{
			// networkLight(),
			// vm(),
			// network("network17"),
			// containerLight(),
			// containerLight2(),
			// // zdbLight(1 * gridtypes.Megabyte),
			// // disk("zmount", 7*gridtypes.Gigabyte),
			// // data(),
			zdb("zdb32", 1*gridtypes.Gigabyte),
			// // container(map[string]string{"volume": "/volume"}),
			// //publicip("pub"),
			// vm(map[string]string{
			// 	"root": "/",
			// }, ""),
			// //containerPublic("pub"),
			// //gateway("http://10.1.1.3", &net),
			// // {
			// // 	Version: Version,
			// // 	Name:    "container_logs",
			// // 	Type:    zos.ZLogsType,
			// // 	Data: gridtypes.MustMarshal(zos.ZLogs{
			// // 		ZMachine: "container",
			// // 		Output:   "ws://185.206.122.21:8080/logs/my-other-container",
			// // 	}),
			// // },
			// // {
			// // 	Version: Version,
			// // 	Name:    "volume",
			// // 	Type:    zos.VolumeType,
			// // 	Data: gridtypes.MustMarshal(zos.Volume{
			// // 		Size: 2 * gridtypes.Gigabyte,
			// // 	}),
			// // },
			// // {
			// // 	Version: Version,
			// // 	Name:    "vm_logs",
			// // 	Type:    zos.ZLogsType,
			// // 	Data: gridtypes.MustMarshal(zos.ZLogs{
			// // 		ZMachine: "vm",
			// // 		Output:   "ws://192.168.123.1:8080/logs/my-vm",
			// 	}),
			// },
			// updated hence version is 1
			// publicip(),
			// containerPublic(),
		},
		SignatureRequirement: gridtypes.SignatureRequirement{
			WeightRequired: 1,
			Requests: []gridtypes.SignatureRequest{
				{
					TwinID: Twin,
					Weight: 1,
				},
			},
		},
	}

	// dl = gridtypes.Deployment{
	// 	Version: 3,
	// 	TwinID:  Twin, //LocalTwin,
	// 	// this contract id must match the one on substrate
	// 	Workloads: []gridtypes.Workload{
	// 		{
	// 			Version: 3,
	// 			Name:    "zdb",
	// 			Type:    zos.ZDBType,
	// 			Data: gridtypes.MustMarshal(zos.ZDB{
	// 				Size:     5 * gridtypes.Gigabyte,
	// 				Mode:     zos.ZDBModeUser,
	// 				Password: "password",
	// 			}),
	// 		},
	// 	},
	// 	SignatureRequirement: gridtypes.SignatureRequirement{
	// 		WeightRequired: 1,
	// 		Requests: []gridtypes.SignatureRequest{
	// 			{
	// 				TwinID: Twin,
	// 				Weight: 1,
	// 			},
	// 		},
	// 	},
	// }

	if err := dl.Valid(); err != nil {
		panic("invalid: " + err.Error())
	}
	// return
	if err := dl.Sign(Twin, identity); err != nil {
		panic(err)
	}

	hash, err := dl.ChallengeHash()
	if err != nil {
		panic("failed to create hash")
	}

	hashHex := hex.EncodeToString(hash)
	fmt.Printf("hash: %s\n", hashHex)
	// create contract

	mgr := substrate.NewManager(Substrate)
	sub, err := mgr.Substrate()
	if err != nil {
		panic(err)
	}

	// if err := sub.AcceptTermsAndConditions(identity, "", ""); err != nil {
	// 	panic(err)
	// }

	// cl, err := rmb.NewRMBClient("redis://localhost:6379")
	// if err != nil {
	// 	panic(err)
	// }
	cl, err := peer.NewRpcClient(context.Background(), Mnemonic, mgr, peer.WithRelay("wss://relay.dev.grid.tf"))
	if err != nil {
		panic(err)
	}

	nodeInfo, err := sub.GetNode(NodeID)
	if err != nil {
		panic(err)
	}
	// node := client.NewNodeClient(uint32(nodeInfo.TwinID), cl)
	fmt.Println("node twin id :", nodeInfo.TwinID)
	node := client.NewNodeClient(uint32(nodeInfo.TwinID), cl)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")

	// stats, err := node.Counters(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("--counters--")
	// enc.Encode(stats)

	// pools, err := node.Pools(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("---pools---")
	// enc.Encode(pools)

	// fmt.Println("---version---")
	// ver, err := node.SystemVersion(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// enc.Encode(ver)

	// return
	// fmt.Println("---gpus---")
	// gpus, err := node.GPUs(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// enc.Encode(gpus)

	countIps := func(dl *gridtypes.Deployment) uint32 {
		var count uint32
		for _, wl := range dl.Workloads {
			if wl.Type == zos.PublicIPType {
				var data zos.PublicIP
				if err := json.Unmarshal(wl.Data, &data); err != nil {
					panic(err)
				}

				if data.V4 {
					count += 1
				}
			}
		}
		return count
	}

	dl.ContractID = 173067
	if false {
		// create
		// if true {
		// policy := substrate.WithCapacityPolicy(cap.AsResources())
		// contractID, err := sub.CreateCapacityReservationContract(identity, 1, policy, nil)
		// if err != nil {
		// 	panic(err)
		// }
		// sub.CancelContract(identity, 50446)
		// return
		ips := countIps(&dl)
		fmt.Println("Required Public Ips: ", ips)
		fmt.Printf("Hash: %s\n", hashHex)

		// contractId, err := sub.UpdateNodeContract(identity, 107732, "", hashHex)
		// if err != nil {
		// 	panic(err)
		// }

		contractId, err := sub.CreateNodeContract(identity, NodeID, "", hashHex, ips, nil)
		if err != nil {
			panic(err)
		}
		dl.ContractID = contractId // from substrate
		// c, err := sub.GetContract(50444)
		// if err != nil {
		// 	panic(err)
		// }
		// if c.ContractType.NodeContract.DeploymentHash.String() != hex.EncodeToString(hash) {
		// 	fmt.Println(hex.EncodeToString(hash))
		// 	fmt.Println(c.ContractType.NodeContract.DeploymentHash.String())
		// 	panic("how")
		// }
		// fmt.Println("ContractID:", contractId)
		err = node.DeploymentDeploy(ctx, dl)
		if err != nil {
			panic(err)
		}
	}
	// dl.ContractID = 137342
	fmt.Println("getting deployment details, contactID: ", dl.ContractID)
	got := gridtypes.Deployment{}
	for {
		got, err = node.DeploymentGet(ctx, dl.ContractID)
		if err == nil {
			break
		}
		fmt.Println("failed to get deployment ", err)
		time.Sleep(2 * time.Second)
	}

	_ = enc.Encode(got)
	fmt.Println(dl.ContractID)
}
