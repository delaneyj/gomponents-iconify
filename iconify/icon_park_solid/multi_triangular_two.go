package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiTriangularTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMultiTriangularTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 41h40L24 6L4 41Z"/><path fill="#fff" d="M34 23.5L24 41L14 23.5h20Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMultiTriangularTwo0)"/>`),
		g.Group(children),
	)
}