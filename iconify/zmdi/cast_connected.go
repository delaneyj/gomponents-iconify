package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 320q27 0 45.5 18.5T64 384H0v-64zm0-85q62 0 105.5 43.5T149 384h-42q0-44-31.5-75.5T0 277v-42zM384 85v214H264q-21-64-68-111T85 120V85h299zM0 149q97 0 166 69t69 166h-43q0-80-56-136T0 192v-43zM427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H277v-43h150V43H43v64H0V43q0-18 12.5-30.5T43 0h384z"/>`),
		g.Group(children),
	)
}