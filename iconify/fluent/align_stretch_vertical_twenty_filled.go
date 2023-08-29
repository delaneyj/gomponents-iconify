package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignStretchVerticalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 3a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1h-15Zm0 13a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1h-15ZM4 13a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v6Zm9 2a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-1Z"/>`),
		g.Group(children),
	)
}