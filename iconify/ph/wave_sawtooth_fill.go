package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaveSawtoothFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm-11.56 94.66l-72 48A8 8 0 0 1 128 184a8 8 0 0 1-8-8V95l-59.56 39.66a8 8 0 1 1-8.88-13.32l72-48A8 8 0 0 1 136 80v81.05l59.56-39.71a8 8 0 0 1 8.88 13.32Z"/>`),
		g.Group(children),
	)
}