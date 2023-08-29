package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRedoSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 2.5a.5.5 0 0 0-1 0v3.843L8.827 3.172a4 4 0 0 0-5.656 5.656l5.025 5.026a.5.5 0 0 0 .707-.708L3.879 8.121A3 3 0 0 1 8.12 3.88L11.243 7H7.499a.5.5 0 0 0 0 1h4.9a.6.6 0 0 0 .6-.6V2.5Z"/>`),
		g.Group(children),
	)
}