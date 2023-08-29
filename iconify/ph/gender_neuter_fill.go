package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderNeuterFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M167.84 108.35a40 40 0 1 1-36.19-36.19a40 40 0 0 1 36.19 36.19ZM216 40v176a16 16 0 0 1-16 16H56a16 16 0 0 1-16-16V40a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16Zm-32 72a56 56 0 1 0-64 55.42v32.31a8.18 8.18 0 0 0 7.47 8.25a8 8 0 0 0 8.53-8v-32.56A56.09 56.09 0 0 0 184 112Z"/>`),
		g.Group(children),
	)
}