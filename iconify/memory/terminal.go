package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 3v16h-1v1H3v-1H2V3h1V2h16v1h1m-2 3H4v12h14V6M9 9v1h1v1h1v2h-1v1H9v1H8v1H6v-2h1v-1h1v-2H7v-1H6V8h2v1h1m2 7v-2h5v2h-5Z"/>`),
		g.Group(children),
	)
}