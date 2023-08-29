package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityAncientOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.52 4.915c.922-.936 1.482-1.96 1.482-2.915h2c0 .956.56 1.98 1.482 2.916c.91.923 2.037 1.617 2.864 1.924c.337.125.554.16.652.16h1v2h-1v2h4v2h-1v7h1v2H2v-2h1v-7H2v-2h4.002V9H5V7h1c.098 0 .316-.035.654-.16c.827-.308 1.955-1.002 2.865-1.925ZM8.001 9v2H16V9H8.002Zm5.796-2a10.28 10.28 0 0 1-.738-.68A8.744 8.744 0 0 1 12 5.03a8.745 8.745 0 0 1-1.058 1.29c-.236.24-.484.467-.74.681H13.8ZM5 13v7h6.002v-4h2v4H19v-7H5Z"/>`),
		g.Group(children),
	)
}