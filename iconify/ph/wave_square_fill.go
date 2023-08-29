package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaveSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm-8 128a8 8 0 0 1-8 8h-72a8 8 0 0 1-8-8V96H64v32a8 8 0 0 1-16 0V88a8 8 0 0 1 8-8h72a8 8 0 0 1 8 8v72h56v-32a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}