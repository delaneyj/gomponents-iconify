package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Progressbar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 5v6h16V5H0zm15 5H1V6h14v4z"/><path fill="currentColor" d="M2 7h7v2H2V7z"/>`),
		g.Group(children),
	)
}