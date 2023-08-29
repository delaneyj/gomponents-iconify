package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10 1.5h3.5m0 0V5m0-3.5l-4 4m-8 4.5v3.5m0 0H5m-3.5 0l4-4"/>`),
		g.Group(children),
	)
}