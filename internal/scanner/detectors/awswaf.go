package detectors

func AWSWaf() *Detector {
	d := &Detector{
		WAFName: "AWSWaf",
		Vendor:  "Amazon",
	}

	d.Checks = []Check{
		CheckHeader("Server", "awselb/2.0"),
	}

	return d
}
