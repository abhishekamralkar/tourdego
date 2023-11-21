package main

type Instance struct {
	Name      string `json:"name"`
	PrivateIP string `json:"privateip"`
	OS        string `json:"os"`
	Tags      Tags   `json:"tags"`
}

type Tags struct {
	Service string
	Project string
}

func getEc2() (instances []Instance) {

	instance1 := Instance{
		Name:      "instance-1",
		PrivateIP: "192.168.2.1",
		OS:        "Linux",
		Tags: Tags{
			Service: "Web",
			Project: "E-Commerce",
		},
	}

	instance2 := Instance{
		Name:      "instance-2",
		PrivateIP: "192.168.2.2",
		OS:        "Linux",
		Tags: Tags{
			Service: "Web",
			Project: "E-Commerce",
		},
	}

	instances = append(instances, instance1)
	instances = append(instances, instance2)

	return instances

}
