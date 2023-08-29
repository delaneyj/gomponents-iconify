package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.95 2H20a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8.05c.329.033.663.05 1 .05c5.523 0 10-4.477 10-10c0-.337-.017-.671-.05-1ZM15 16v2h2v-2h-2ZM11.938 2A8 8 0 0 1 3 10.938V3a1 1 0 0 1 1-1h7.938Z"/>`),
		g.Group(children),
	)
}