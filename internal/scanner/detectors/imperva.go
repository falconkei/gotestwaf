package detectors

func SecureSphere() *Detector {
	d := &Detector{
		WAFName: "SecureSphere",
		Vendor:  "Imperva Inc.",
	}

	d.Checks = []Check{
		CheckContent("<(title|h2)>Error"),
		CheckContent("The incident ID is"),
		CheckContent("This page can't be displayed"),
		CheckContent("Contact support for additional information"),
	}

	return d
}

func SecureSphereCustom() *Detector {
	d := &Detector{
		WAFName: "SecureSphereCustom",
		Vendor:  "Imperva Inc.",
	}

	d.Checks = []Check{
		CheckStatusCode(406),
		CheckContent("If you keep seeing this page, please contact support."),
		CheckContent("Please provide your company name and other contact information"),
	}

	return d
}

func Incapsula() *Detector {
	d := &Detector{
		WAFName: "Incapsula",
		Vendor:  "Imperva Inc.",
	}

	d.Checks = []Check{
		CheckStatusCode(403),
		CheckContent("Incapsula incident ID"),
		CheckContent("/_Incapsula_Resource"),
	}

	return d
}
