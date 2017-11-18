package sitemap

import "bytes"

// Sitemap represents a single sitemap.
type Sitemap struct {
	sites []*Site
}

// New creates a new sitemap.
func New() *Sitemap {
	return &Sitemap{}
}

// Add adds a site to the sitemap.
func (sitemap *Sitemap) Add(url string) {
	sitemap.sites = append(sitemap.sites, &Site{
		URL: url,
	})
}

// Text creates a list of your URLs in text format.
func (sitemap *Sitemap) Text() string {
	var buffer bytes.Buffer

	for _, site := range sitemap.sites {
		buffer.WriteString(site.URL)
		buffer.WriteByte('\n')
	}

	return buffer.String()
}
