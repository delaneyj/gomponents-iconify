package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectCurveThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21c-4.97 0-9-1.79-9-4s4.03-4 9-4s9 1.79 9 4s-4.03 4-9 4Zm0-19a3 3 0 0 1 3 3v1H9V5a3 3 0 0 1 3-3ZM3.5 15.5l4-7m13 7l-4-7"/>`),
		g.Group(children),
	)
}