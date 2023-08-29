package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassWhiskey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m4.818 7l3.334 20h15.696L27.18 7H4.818zm2.364 2H24.82l-1.666 10H10.67l.33 2h11.82l-.666 4H9.848L7.182 9z"/>`),
		g.Group(children),
	)
}