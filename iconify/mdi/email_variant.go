package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13L2 6.76V6c0-1.11.89-2 2-2h16a2 2 0 0 1 2 2v.75L12 13m10 5a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9.11l2 1.25V18h16v-7.64l2-1.25V18Z"/>`),
		g.Group(children),
	)
}