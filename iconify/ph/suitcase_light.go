package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 58h-42V48a22 22 0 0 0-22-22h-48a22 22 0 0 0-22 22v10H40a14 14 0 0 0-14 14v128a14 14 0 0 0 14 14h176a14 14 0 0 0 14-14V72a14 14 0 0 0-14-14ZM94 48a10 10 0 0 1 10-10h48a10 10 0 0 1 10 10v10H94Zm68 22v132H94V70ZM38 200V72a2 2 0 0 1 2-2h42v132H40a2 2 0 0 1-2-2Zm180 0a2 2 0 0 1-2 2h-42V70h42a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}