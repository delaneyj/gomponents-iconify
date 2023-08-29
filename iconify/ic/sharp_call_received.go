package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCallReceived(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5.41L18.59 4L7 15.59V9H5v10h10v-2H8.41L20 5.41z"/>`),
		g.Group(children),
	)
}