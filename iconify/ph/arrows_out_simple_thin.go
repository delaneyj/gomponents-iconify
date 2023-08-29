package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 48v48a4 4 0 0 1-8 0V57.66l-57.17 57.17a4 4 0 0 1-5.66-5.66L198.34 52H160a4 4 0 0 1 0-8h48a4 4 0 0 1 4 4Zm-102.83 93.17L52 198.34V160a4 4 0 0 0-8 0v48a4 4 0 0 0 4 4h48a4 4 0 0 0 0-8H57.66l57.17-57.17a4 4 0 0 0-5.66-5.66Z"/>`),
		g.Group(children),
	)
}