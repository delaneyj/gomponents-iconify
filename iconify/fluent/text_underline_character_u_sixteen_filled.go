package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineCharacterUSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2.75a.75.75 0 0 0-1.5 0V7.5a3.5 3.5 0 1 0 7 0V2.75a.75.75 0 0 0-1.5 0V7.5a2 2 0 1 1-4 0V2.75ZM4.75 12.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5h-6.5Z"/>`),
		g.Group(children),
	)
}