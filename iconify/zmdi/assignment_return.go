package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignmentReturn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h89q7-19 23.5-31T192 5t36.5 12T252 48h89zm-149 0q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5zm85 256v-85h-85v-64L85 261l107 107v-64h85z"/>`),
		g.Group(children),
	)
}