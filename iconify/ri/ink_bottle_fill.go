package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkBottleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 8.997l4.371 1.748a1 1 0 0 1 .629.929v9.323a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-9.323a1 1 0 0 1 .629-.929L8 8.997h8Zm4 5H8v5h12v-5Zm-4-11a1 1 0 0 1 1 1v4H7v-4a1 1 0 0 1 1-1h8Z"/>`),
		g.Group(children),
	)
}