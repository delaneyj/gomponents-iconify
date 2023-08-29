package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdentificationCardLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198 112a6 6 0 0 1-6 6h-40a6 6 0 0 1 0-12h40a6 6 0 0 1 6 6Zm-6 26h-40a6 6 0 0 0 0 12h40a6 6 0 0 0 0-12Zm38-82v144a14 14 0 0 1-14 14H40a14 14 0 0 1-14-14V56a14 14 0 0 1 14-14h176a14 14 0 0 1 14 14Zm-12 0a2 2 0 0 0-2-2H40a2 2 0 0 0-2 2v144a2 2 0 0 0 2 2h176a2 2 0 0 0 2-2Zm-84.19 110.51a6 6 0 0 1-11.62 3C119.34 158.38 108.08 150 96 150s-23.33 8.38-26.19 19.5a6 6 0 0 1-11.62-3A38 38 0 0 1 76.78 143a30 30 0 1 1 38.45 0a38 38 0 0 1 18.58 23.51ZM96 138a18 18 0 1 0-18-18a18 18 0 0 0 18 18Z"/>`),
		g.Group(children),
	)
}