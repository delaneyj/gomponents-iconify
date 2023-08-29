package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DismissThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M26.29 4.293a1 1 0 1 1 1.414 1.414L17.413 15.998L27.704 26.29a1 1 0 1 1-1.414 1.414L16 17.413L5.707 27.704a1 1 0 0 1-1.414-1.414L14.584 16L4.293 5.707a1 1 0 0 1 1.414-1.414L16 14.584L26.29 4.293Z"/>`),
		g.Group(children),
	)
}