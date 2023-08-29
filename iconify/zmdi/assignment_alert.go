package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignmentAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h89q7-19 23.5-31T192 5t36.5 12T252 48h89zM213 368v-43h-42v43h42zm0-85V155h-42v128h42zM192 91q9 0 15-6.5t6-15t-6-15t-15-6.5t-15 6.5t-6 15t6 15t15 6.5z"/>`),
		g.Group(children),
	)
}