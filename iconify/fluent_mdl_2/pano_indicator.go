package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoIndicator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1472 896q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19zM0 279q241 115 499 172t525 58q266 0 524-57t500-173v1488q-241-115-499-173t-525-58q-266 0-524 58T0 1767V279zm128 1298q153-59 311-96t321-54l-312-312l-320 319v143zm802-161q24-1 48-1t49-1q197 0 391 30l-330-329l-229 229l71 72zm990 18l-192-191l-166 165l83 83q70 18 138 39t137 47v-143zm0-965q-215 83-440 127t-456 44q-231 0-456-44T128 469v785l320-321l320 321l320-321l384 384l256-256l192 193V469z"/>`),
		g.Group(children),
	)
}