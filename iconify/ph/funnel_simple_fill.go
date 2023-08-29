package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm-72 136h-32a8 8 0 0 1 0-16h32a8 8 0 0 1 0 16Zm32-40H80a8 8 0 0 1 0-16h96a8 8 0 0 1 0 16Zm32-40H48a8 8 0 0 1 0-16h160a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}