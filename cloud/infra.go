package main

type Instance struct {
	Name      string
	PrivateIP string
	OS        string
}

func getEc2() (instances []Instance) {

	instance1 := Instance{
		Name:      "instance-1",
		PrivateIP: "192.168.2.1",
		OS:        "Linux",
	}

	instance2 := Instance{
		Name:      "instance-2",
		PrivateIP: "192.168.2.2",
		OS:        "Linux",
	}

	instances = append(instances, instance1)
	instances = append(instances, instance2)

	return instances

}
