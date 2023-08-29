package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListMagnifyingGlassLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M34 64a6 6 0 0 1 6-6h176a6 6 0 0 1 0 12H40a6 6 0 0 1-6-6Zm6 70h72a6 6 0 0 0 0-12H40a6 6 0 0 0 0 12Zm88 52H40a6 6 0 0 0 0 12h88a6 6 0 0 0 0-12Zm108.24 10.24a6 6 0 0 1-8.48 0l-21.49-21.48a38.06 38.06 0 1 1 8.49-8.49l21.48 21.49a6 6 0 0 1 0 8.48ZM184 170a26 26 0 1 0-26-26a26 26 0 0 0 26 26Z"/>`),
		g.Group(children),
	)
}