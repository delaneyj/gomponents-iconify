package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSatellite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3H3v18h18V3zM5 4.99h3C8 6.65 6.66 8 5 8V4.99zM5 12v-2c2.76 0 5-2.25 5-5.01h2C12 8.86 8.87 12 5 12zm0 6l3.5-4.5l2.5 3.01L14.5 12l4.5 6H5z"/>`),
		g.Group(children),
	)
}