package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleDownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M210.83 125.17a4 4 0 0 1 0 5.66l-80 80a4 4 0 0 1-5.66 0l-80-80a4 4 0 1 1 5.66-5.66L128 202.34l77.17-77.17a4 4 0 0 1 5.66 0Zm-85.66 5.66a4 4 0 0 0 5.66 0l80-80a4 4 0 1 0-5.66-5.66L128 122.34L50.83 45.17a4 4 0 0 0-5.66 5.66Z"/>`),
		g.Group(children),
	)
}