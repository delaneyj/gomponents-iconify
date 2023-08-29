package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWindyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 18v-2h3.5a3.5 3.5 0 1 0-2.5-5.95V10a6 6 0 0 0-12 0v.007H1V10a8 8 0 0 1 15.458-2.901A5.5 5.5 0 1 1 17.5 18H14Zm-8 2h10v2H6v-2Zm0-8h8v2H6v-2Zm-4 4h10v2H2v-2Z"/>`),
		g.Group(children),
	)
}