package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUpDownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M178.83 173.17a4 4 0 0 1 0 5.66l-48 48a4 4 0 0 1-5.66 0l-48-48a4 4 0 0 1 5.66-5.66L128 218.34l45.17-45.17a4 4 0 0 1 5.66 0Zm-96-90.34L128 37.66l45.17 45.17a4 4 0 1 0 5.66-5.66l-48-48a4 4 0 0 0-5.66 0l-48 48a4 4 0 0 0 5.66 5.66Z"/>`),
		g.Group(children),
	)
}