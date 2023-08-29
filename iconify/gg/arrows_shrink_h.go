package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsShrinkH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 7h2v10H1V7Zm7.448.757l1.414 1.415L8.033 11h7.933l-1.828-1.828l1.414-1.415L19.795 12l-4.243 4.243l-1.414-1.415L15.966 13H8.034l1.828 1.828l-1.414 1.415L4.205 12l4.243-4.243ZM23 7h-2v10h2V7Z"/>`),
		g.Group(children),
	)
}