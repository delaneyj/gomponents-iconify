package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 14C2 7.373 7.373 2 14 2s12 5.373 12 12s-5.373 12-12 12S2 20.627 2 14Zm14.5-4.5a.75.75 0 0 0 0-1.5H12a.75.75 0 0 0-.75.709l-.25 4.5a.75.75 0 0 0 .75.791h.318l.739.004c.55.005 1.122.013 1.354.027a2.75 2.75 0 1 1-2.771 4.127a.75.75 0 1 0-1.288.769a4.25 4.25 0 1 0 4.19-6.39a34.16 34.16 0 0 0-1.474-.033l-.276-.001l.167-3.003h3.793Z"/>`),
		g.Group(children),
	)
}