package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterHorizontalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 0v3H3.5v4H7v1H1v4h6v3h1v-3h6V8H8V7h3.5V3H8V0H7Z"/>`),
		g.Group(children),
	)
}