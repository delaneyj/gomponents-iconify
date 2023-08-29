package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NTwoSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298zm-85 171v-43q0-18-12.5-30.5T213 85h-85v43h85v43h-42q-18 0-30.5 12.5T128 213v86h128v-43h-85v-43h42q18 0 30.5-12.5T256 171z"/>`),
		g.Group(children),
	)
}