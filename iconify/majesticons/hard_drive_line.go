package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15h14M5 15v4a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-4M5 15V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10m-6 3h3"/>`),
		g.Group(children),
	)
}