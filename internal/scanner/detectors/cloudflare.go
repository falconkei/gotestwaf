package detectors

func Cloudflare() *Detector {
	d := &Detector{
		WAFName: "Cloudflare",
		Vendor:  "Cloudflare Inc",
	}

	d.Checks = []Check{
		CheckHeader("Server", "cloudflare"),
		CheckContent("Ray ID"),
	}

	return d
}
