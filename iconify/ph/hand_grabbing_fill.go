package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandGrabbingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 104v48a88 88 0 0 1-176 0v-16a16 16 0 0 1 32 0v8a8 8 0 0 0 16 0V88a16 16 0 0 1 32 0v16a8 8 0 0 0 16 0V88a16 16 0 0 1 32 0v16a8 8 0 0 0 16 0a16 16 0 0 1 32 0Z"/>`),
		g.Group(children),
	)
}