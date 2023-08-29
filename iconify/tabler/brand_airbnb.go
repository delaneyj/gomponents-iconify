package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAirbnb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10c-2 0-3 1-3 3c0 1.5 1.494 3.535 3 5.5c1 1 1.5 1.5 2.5 2s2.5 1 4.5-.5s1.5-3.5.5-6s-2.333-5.5-5-9.5C13.666 3.5 13 3 11.997 3c-1 0-1.623.45-2.497 1.5c-2.667 4-4 7-5 9.5S3 18.5 5 20s3.5 1 4.5.5s1.5-1 2.5-2c1.506-1.965 3-4 3-5.5c0-2-1-3-3-3z"/>`),
		g.Group(children),
	)
}