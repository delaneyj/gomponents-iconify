package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func History(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 1a7 7 0 1 1 0 14v-1.5c1.469 0 2.85-.572 3.889-1.611S15.5 9.469 15.5 8c0-1.469-.572-2.85-1.611-3.889S11.469 2.5 10 2.5c-1.469 0-2.85.572-3.889 1.611A5.455 5.455 0 0 0 4.591 7H7.5L4 11L.5 7h2.571A7.001 7.001 0 0 1 10 1zm3 6v2H9V4h2v3z"/>`),
		g.Group(children),
	)
}