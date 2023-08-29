package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayListTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 18v2H2v-2h20ZM2 3.5l8 5l-8 5v-10ZM22 11v2H12v-2h10Zm0-7v2H12V4h10Z"/>`),
		g.Group(children),
	)
}