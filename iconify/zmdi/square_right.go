package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v299q0 17-12.5 29.5T427 384H43q-18 0-30.5-12.5T0 342v-86h43v86h384V42H43v86H0V43q0-18 12.5-30.5T43 0h384zM213 277v-64H0v-42h213v-64l86 85z"/>`),
		g.Group(children),
	)
}