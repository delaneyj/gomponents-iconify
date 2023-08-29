package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignmentAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h89q7-19 23.5-31T192 5t36.5 12T252 48h89zm-149 0q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5zm0 85q-27 0-45.5 19T128 197.5t18.5 45T192 261t45.5-18.5t18.5-45t-18.5-45.5t-45.5-19zm128 256v-30q0-19-23.5-35T244 300.5t-52-7.5t-52 7.5T87.5 324T64 359v30h256z"/>`),
		g.Group(children),
	)
}