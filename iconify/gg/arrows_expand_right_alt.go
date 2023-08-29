package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandRightAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.9 2.1v2h4.585l-5.778 5.78l1.414 1.414L19.9 5.515V10.1h2v-8h-8ZM5.515 19.9H10.1v2h-8v-8h2v4.585l5.778-5.778l1.414 1.414L5.515 19.9ZM9.172 7.757L7.757 9.172l7.071 7.07l1.415-1.414l-7.071-7.07Z"/>`),
		g.Group(children),
	)
}