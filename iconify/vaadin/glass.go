package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 15H9V7l6-7H0l6 7v8H4c-2 0-2 1-2 1h11s0-1-2-1zm1.9-14l-1.8 2H3.9L2.2 1h10.7zM7 15V7h1v8H7z"/>`),
		g.Group(children),
	)
}