package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.965 1.6l.035.034l-8.832 8.855H24v3.556H7.166L16 22.9l-.035.034h-5.319L0 12.268L10.647 1.601z"/>`),
		g.Group(children),
	)
}