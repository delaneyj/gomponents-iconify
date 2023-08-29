package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaggingTaxi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24.798 2l-4.442 6.664A2.994 2.994 0 0 1 17.858 10H12a3.003 3.003 0 0 0-3 3v6a2.002 2.002 0 0 0 2 2v7a2.002 2.002 0 0 0 2 2h4a2.002 2.002 0 0 0 2-2V16h-2v12h-4v-9h-2v-6a1 1 0 0 1 1-1h5.86a4.988 4.988 0 0 0 4.16-2.227l4.442-6.664Z"/><path fill="currentColor" d="M11 5a4 4 0 1 1 4 4a4.004 4.004 0 0 1-4-4Zm2 0a2 2 0 1 0 2-2a2.002 2.002 0 0 0-2 2Z"/>`),
		g.Group(children),
	)
}