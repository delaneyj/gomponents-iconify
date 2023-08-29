package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFilters(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#F57C00" d="M29 23H19L7 9h34z"/><path fill="#FF9800" d="m29 38l-10 6V23h10zM41.5 9h-35C5.7 9 5 8.3 5 7.5S5.7 6 6.5 6h35c.8 0 1.5.7 1.5 1.5S42.3 9 41.5 9z"/><circle cx="38" cy="38" r="10" fill="#F44336"/><path fill="#fff" d="M32 36h12v4H32z"/>`),
		g.Group(children),
	)
}