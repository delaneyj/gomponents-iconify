package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckSquareOffsetThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 48v160a12 12 0 0 1-12 12h-72a4 4 0 0 1 0-8h72a4 4 0 0 0 4-4V48a4 4 0 0 0-4-4H48a4 4 0 0 0-4 4v96a4 4 0 0 1-8 0V48a12 12 0 0 1 12-12h160a12 12 0 0 1 12 12ZM117.17 157.17L64 210.34l-21.17-21.17a4 4 0 0 0-5.66 5.66l24 24a4 4 0 0 0 5.66 0l56-56a4 4 0 0 0-5.66-5.66Z"/>`),
		g.Group(children),
	)
}