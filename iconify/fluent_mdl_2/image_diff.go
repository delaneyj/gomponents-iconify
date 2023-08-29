package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageDiff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v1920H0V0h2048zM128 896h933L576 410L128 859v37zm1573 0l-293-293l-230 229l65 64h458zM128 1792h933l-485-486l-448 449v37zm1115 0h458l-293-293l-230 229l65 64zm677-768H128v549l448-447l512 512l320-321l475 475h37v-768zm0-128V128H128v549l448-447l512 512l320-321l475 475h37z"/>`),
		g.Group(children),
	)
}