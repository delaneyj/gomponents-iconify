package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageConveyorBelt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8 3.5v11m8-11v11m-1-11a3 3 0 1 0-6 0M4 21h1m4 0h1m4 0h1m4 0h1m0-6.5H4v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L4 3.676V3.5h16v.176l-.203 1.346a26.749 26.749 0 0 0 0 7.956L20 14.323v.177Zm1 4H3a2.5 2.5 0 0 0 0 5h18a2.5 2.5 0 0 0 0-5Z"/>`),
		g.Group(children),
	)
}