package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 15V1L9 8zm-7 0V1L2 8zM0 1h2v14H0V1z"/>`),
		g.Group(children),
	)
}