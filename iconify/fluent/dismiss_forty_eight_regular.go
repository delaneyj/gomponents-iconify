package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DismissFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.384 6.616a1.25 1.25 0 1 0-1.768 1.768L22.232 24L6.616 39.616a1.25 1.25 0 0 0 1.768 1.768L24 25.768l15.615 15.615a1.25 1.25 0 1 0 1.768-1.768L25.768 24L41.383 8.384a1.25 1.25 0 1 0-1.767-1.767L24 22.232L8.384 6.616Z"/>`),
		g.Group(children),
	)
}