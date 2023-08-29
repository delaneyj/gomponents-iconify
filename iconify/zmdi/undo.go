package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M224 107q73 0 131.5 43T437 261l-51 16q-17-51-61.5-84T224 160q-61 0-109 40l77 77H0V85l77 77q64-55 147-55z"/>`),
		g.Group(children),
	)
}