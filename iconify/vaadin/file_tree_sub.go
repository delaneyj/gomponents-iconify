package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTreeSub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 11v1H7v-2h5V6H4v1H3V5h6V1H0v4h2v3h2v2h2v3h2v2h8v-4z"/>`),
		g.Group(children),
	)
}