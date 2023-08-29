package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 2a.5.5 0 0 0 0 1h3.793L3 12.293V8.5a.5.5 0 0 0-1 0v4.9a.6.6 0 0 0 .6.6h4.9a.5.5 0 0 0 0-1H3.707L13 3.707V7.5a.5.5 0 0 0 1 0V2.6a.6.6 0 0 0-.6-.6H8.5Z"/>`),
		g.Group(children),
	)
}