package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 4h9v1h2v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-2v1H8v1H4v1H1v-2h2v-1h1v-2H3v-1H2V8h1V7h1V6h1V5h2V4m10 4V7h-2V6H8v1H6v1H5v1H4v4h1v1h1v1h2v1h7v-1h2v-1h1v-1h1V9h-1V8h-1Z"/>`),
		g.Group(children),
	)
}