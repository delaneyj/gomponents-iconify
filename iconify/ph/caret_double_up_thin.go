package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M210.83 205.17a4 4 0 0 1-5.66 5.66L128 133.66l-77.17 77.17a4 4 0 0 1-5.66-5.66l80-80a4 4 0 0 1 5.66 0Zm-160-74.34L128 53.66l77.17 77.17a4 4 0 0 0 5.66-5.66l-80-80a4 4 0 0 0-5.66 0l-80 80a4 4 0 0 0 5.66 5.66Z"/>`),
		g.Group(children),
	)
}