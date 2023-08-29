package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainClimbing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1408 817l616 1231H24L768 561V0h512v384H896v177l320 640l192-384zM896 128v128h256V128H896zm93 905L832 719l-157 314l157 156l157-156zm443 887l-383-767l-217 218l-217-218l-383 767h1200zm144 0h240l-408-817l-120 241l288 576z"/>`),
		g.Group(children),
	)
}