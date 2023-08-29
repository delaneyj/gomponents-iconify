package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m138.83 130.83l-80 80a4 4 0 0 1-5.66-5.66L130.34 128L53.17 50.83a4 4 0 0 1 5.66-5.66l80 80a4 4 0 0 1 0 5.66Zm80-5.66l-80-80a4 4 0 0 0-5.66 5.66L210.34 128l-77.17 77.17a4 4 0 0 0 5.66 5.66l80-80a4 4 0 0 0 0-5.66Z"/>`),
		g.Group(children),
	)
}