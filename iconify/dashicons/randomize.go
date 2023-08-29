package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Randomize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 6.01L14 9V7h-4l-5 8H2v-2h2l5-8h5V3zM2 5h3l1.15 2.17l-1.12 1.8L4 7H2V5zm16 9.01L14 17v-2H9l-1.15-2.17l1.12-1.8L10 13h4v-2z"/>`),
		g.Group(children),
	)
}