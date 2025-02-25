package builder

import (
	_ "embed"
	"fmt"
	"github.com/falcosecurity/driverkit/pkg/kernelrelease"
)

//go:embed templates/rocky.sh
var rockyTemplate string

// TargetTypeRocky identifies the Rocky target.
const TargetTypeRocky Type = "rocky"

func init() {
	BuilderByTarget[TargetTypeRocky] = &rocky{}
}

type rockyTemplateData struct {
	commonTemplateData
	KernelDownloadURL string
}

// rocky is a driverkit target.
type rocky struct {
}

func (c *rocky) Name() string {
	return TargetTypeRocky.String()
}

func (c *rocky) TemplateScript() string {
	return rockyTemplate
}

func (c *rocky) URLs(_ Config, kr kernelrelease.KernelRelease) ([]string, error) {
	return fetchRockyKernelURLS(kr), nil
}

func (c *rocky) TemplateData(cfg Config, kr kernelrelease.KernelRelease, urls []string) interface{} {
	return rockyTemplateData{
		commonTemplateData: cfg.toTemplateData(c, kr),
		KernelDownloadURL:  urls[0],
	}
}

func fetchRockyKernelURLS(kr kernelrelease.KernelRelease) []string {
	rockyReleases := []string{
		"8",
		"8.7",
		"9",
		"9.1",
	}

	rockyVaultReleases := []string{
		"8.3",
		"8.4",
		"8.5",
		"8.6",
		"9.1",
	}

	urls := []string{}
	for _, r := range rockyReleases {
		if r >= "9" {
			urls = append(urls, fmt.Sprintf(
				"https://download.rockylinux.org/pub/rocky/%s/AppStream/%s/os/Packages/k/kernel-devel-%s%s.rpm",
				r,
				kr.Architecture.ToNonDeb(),
				kr.Fullversion,
				kr.FullExtraversion,
			))
		} else {
			urls = append(urls, fmt.Sprintf(
				"https://download.rockylinux.org/pub/rocky/%s/BaseOS/%s/os/Packages/k/kernel-devel-%s%s.rpm",
				r,
				kr.Architecture.ToNonDeb(),
				kr.Fullversion,
				kr.FullExtraversion,
			))
		}
	}
	for _, r := range rockyVaultReleases {
		if r >= "9" {
			urls = append(urls, fmt.Sprintf(
				"https://download.rockylinux.org/vault/rocky/%s/AppStream/%s/os/Packages/k/kernel-devel-%s%s.rpm",
				r,
				kr.Architecture.ToNonDeb(),
				kr.Fullversion,
				kr.FullExtraversion,
			))
		} else {
			urls = append(urls, fmt.Sprintf(
				"https://download.rockylinux.org/vault/rocky/%s/BaseOS/%s/os/Packages/k/kernel-devel-%s%s.rpm",
				r,
				kr.Architecture.ToNonDeb(),
				kr.Fullversion,
				kr.FullExtraversion,
			))
		}
	}
	return urls
}
