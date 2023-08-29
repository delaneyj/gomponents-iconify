package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toolbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1V6m7 0h4V4H9v2m10 2H3v4h3v-2h3v2h4v-2h3v2h3V8M3 18h16v-4h-3v2h-3v-2H9v2H6v-2H3v4Z"/>`),
		g.Group(children),
	)
}