package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 17h-9v-2h9v2Zm-11 0H2v-2h9v2Zm4-4H2v-2h13v2Zm0-4H2V7h13v2Z"/>`),
		g.Group(children),
	)
}