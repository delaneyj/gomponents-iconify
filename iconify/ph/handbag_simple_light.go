package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandbagSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m237.9 198.36l-14.25-120a14.06 14.06 0 0 0-14-12.36H174v-2a46 46 0 0 0-92 0v2H46.33a14.06 14.06 0 0 0-14 12.36l-14.25 120a14 14 0 0 0 14 15.64h191.84a14 14 0 0 0 14-15.64ZM94 64a34 34 0 0 1 68 0v2H94Zm131.5 137.3a2.07 2.07 0 0 1-1.58.7H32.08a2.07 2.07 0 0 1-1.58-.7a1.92 1.92 0 0 1-.49-1.53l14.26-120A2 2 0 0 1 46.33 78h163.34a2 2 0 0 1 2.06 1.77l14.26 120a1.92 1.92 0 0 1-.49 1.53Z"/>`),
		g.Group(children),
	)
}