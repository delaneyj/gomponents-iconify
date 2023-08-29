package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditStraight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4a7 7 0 0 1 7 7H5a7 7 0 0 1 7-7Zm-7 9H1v-2h4v2Zm14 0a7 7 0 1 1-14 0h14Zm0 0v-2h4v2h-4Z"/>`),
		g.Group(children),
	)
}