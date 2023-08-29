package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorLineAccentTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.004 15.631a2 2 0 0 0 2.571 1.784l4.293-1.288a3.25 3.25 0 0 0 1.159-.627H19.5a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1v-3a1 1 0 0 1 .504-.869Z"/>`),
		g.Group(children),
	)
}