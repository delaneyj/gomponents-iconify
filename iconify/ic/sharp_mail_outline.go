package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpMailOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4H2.01L2 20h20V4zm-2 14H4V8l8 5l8-5v10zm-8-7L4 6h16l-8 5z"/>`),
		g.Group(children),
	)
}