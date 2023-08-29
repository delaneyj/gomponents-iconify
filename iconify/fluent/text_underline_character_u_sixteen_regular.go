package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineCharacterUSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 2.5a.5.5 0 0 0-1 0v5a3.5 3.5 0 1 0 7 0v-5a.5.5 0 0 0-1 0v5a2.5 2.5 0 0 1-5 0v-5ZM4.5 13a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1h-7Z"/>`),
		g.Group(children),
	)
}