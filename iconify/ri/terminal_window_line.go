package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalWindowLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 9V5H4v4h16Zm0 2H4v8h16v-8ZM3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm2 9h3v5H5v-5Zm0-6h2v2H5V6Zm4 0h2v2H9V6Z"/>`),
		g.Group(children),
	)
}