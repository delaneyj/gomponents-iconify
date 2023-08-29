package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterHorizontalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 0v3.5m0 8V15m0-8.5v2M4 3.5v3h7v-3H4Zm-2.5 5v3h12v-3h-12Z"/>`),
		g.Group(children),
	)
}