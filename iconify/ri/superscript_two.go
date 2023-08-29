package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperscriptTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7v13H9V7H3V5h12v2h-4Zm8.55-.42a.8.8 0 1 0-1.32-.36l-1.154.33A2.001 2.001 0 1 1 21 6c0 .573-.24 1.09-.627 1.454L18.744 9H21v1h-4V9l2.55-2.42Z"/>`),
		g.Group(children),
	)
}