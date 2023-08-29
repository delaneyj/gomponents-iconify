package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinFiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.657 15.657L12 21.314l-5.657-5.657a8 8 0 1 1 11.314 0ZM5 22h14v2H5v-2Z"/>`),
		g.Group(children),
	)
}