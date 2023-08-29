package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298zm-85 160v-32q0-18-12.5-30.5T213 85h-85v43h85v43h-42v42h42v43h-85v43h85q18 0 30.5-12.5T256 256v-32q0-13-9.5-22.5T224 192q13 0 22.5-9.5T256 160z"/>`),
		g.Group(children),
	)
}