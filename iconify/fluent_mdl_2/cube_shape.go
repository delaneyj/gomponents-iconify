package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeShape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 569v1038l-832 417l-832-417V569l832-417l832 417zM960 296L335 608l625 312l625-312l-625-312zM256 1528l640 321v-817L256 711v817zm1408 0V711l-640 321v817l640-321z"/>`),
		g.Group(children),
	)
}