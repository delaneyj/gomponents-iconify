package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10.5 7.015V.5h3v6.515a6 6 0 0 0 1.757 4.242L17.5 13.5v10h-11v-10l2.243-2.243A6 6 0 0 0 10.5 7.015Z"/>`),
		g.Group(children),
	)
}