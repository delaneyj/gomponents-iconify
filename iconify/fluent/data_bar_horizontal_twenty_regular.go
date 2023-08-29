package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarHorizontalTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4a2 2 0 0 1 2-2h5a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2Zm2-1a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2H4Zm-2 7a2 2 0 0 1 2-2h8a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2Zm2-1a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2H4Zm-2 7a2 2 0 0 1 2-2h12a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2Zm2-1a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2H4Z"/>`),
		g.Group(children),
	)
}