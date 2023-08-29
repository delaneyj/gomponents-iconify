package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IBeam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M39.557 19L323.44 273h149.003L188.56 19H39.558zM25 30.13v25.847L311 311.87v-25.846L25 30.13zm64 107.263v34.584L375 427.87v-84.843l-64-13.002L89 137.393zM77.555 185.89l-42.9 10.723l287.79 257.498l42.9-10.723L77.556 185.89zM25 212.13v23.847L311 491.87v-23.847L25 212.13zM329 291v21.973l64 16v126.054l-64 16V493h158v-21.973l-64-16V328.973l64-16V291H329zm112 135.865v14.108l21.88 5.47L441 426.865z"/>`),
		g.Group(children),
	)
}