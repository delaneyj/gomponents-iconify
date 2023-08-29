package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PulseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 7.539l6 14L18.66 13H23v-2h-5.66L15 16.461l-6-14L5.34 11H1v2h5.66L9 7.539Z"/>`),
		g.Group(children),
	)
}