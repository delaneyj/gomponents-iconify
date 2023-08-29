package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightDoubleRoundUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 22v-3h-1v-2h-1v-2H9v-1H8v-1H7v-1H5v-1H3v-1H0V8h4v1h2v1h2v1h1v1h1v1h1v1h1v2h1v2h1v4h-2M0 12h2v1h3v1h1v1h1v1h1v1h1v3h1v2H8v-1H7v-3H6v-1H5v-1H4v-1H1v-1H0v-2Z"/>`),
		g.Group(children),
	)
}