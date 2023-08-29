package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageProcessing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m2 14h15V3H3v13h1v-1m2-7h2v2H6V8m4 0h2v2h-2V8m4 0h2v2h-2V8Z"/>`),
		g.Group(children),
	)
}