package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseLineFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 208h-16v-92.45a16 16 0 0 0-5.17-11.78l-80-75.48a1.14 1.14 0 0 1-.11-.11a16 16 0 0 0-21.53 0l-.11.11l-79.91 75.48A16 16 0 0 0 32 115.55V208H16a8 8 0 0 0 0 16h224a8 8 0 0 0 0-16Zm-88 0h-48v-48a8 8 0 0 1 8-8h32a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}