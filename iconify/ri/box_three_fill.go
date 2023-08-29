package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.502 5.922L12 1L3.498 5.922L12 10.845l8.502-4.923ZM2.5 7.656V17.5l8.5 4.921v-9.845l-8.5-4.92ZM13 22.42l8.5-4.921V7.655L13 12.576v9.845Z"/>`),
		g.Group(children),
	)
}