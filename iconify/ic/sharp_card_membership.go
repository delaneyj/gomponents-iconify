package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCardMembership(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H2v15h6v5l4-2l4 2v-5h6V2zm-2 13H4v-2h16v2zm0-5H4V4h16v6z"/>`),
		g.Group(children),
	)
}