package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TidalIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#0A0B09" d="m128.004 85.339l42.664 42.67l-42.664 42.667l-42.669-42.667l42.669-42.67ZM42.667.002L85.335 42.67L42.667 85.34L0 42.67L42.667.002Zm170.666 0L256 42.67l-42.667 42.67l-42.666-42.67l-42.663 42.669l-42.669-42.67L128.004 0l42.663 42.665L213.333.002Z"/>`),
		g.Group(children),
	)
}