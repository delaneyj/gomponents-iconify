package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 1H3v6H0v1h3v6h4V8h1v3.5h4V8h3V7h-3V3.5H8V7H7V1Z"/>`),
		g.Group(children),
	)
}