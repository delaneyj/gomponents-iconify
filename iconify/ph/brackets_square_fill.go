package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BracketsSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16ZM104 176a8 8 0 0 1 0 16H72a8 8 0 0 1-8-8V72a8 8 0 0 1 8-8h32a8 8 0 0 1 0 16H80v96Zm88 8a8 8 0 0 1-8 8h-32a8 8 0 0 1 0-16h24V80h-24a8 8 0 0 1 0-16h32a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}