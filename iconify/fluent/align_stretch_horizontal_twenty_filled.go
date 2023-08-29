package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignStretchHorizontalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 17.5a.5.5 0 0 0 1 0v-15a.5.5 0 0 0-1 0v15Zm13 0a.5.5 0 0 0 1 0v-15a.5.5 0 0 0-1 0v15ZM13 16a2 2 0 0 0 2-2v-1a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h6Zm2-9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1Z"/>`),
		g.Group(children),
	)
}