package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewWeek(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 43q9 0 15.5 6t6.5 15v256q0 9-6.5 15T85 341H21q-8 0-14.5-6T0 320V64q0-9 6.5-15T21 43h64zm299 0q9 0 15 6t6 15v256q0 9-6 15t-15 6h-64q-9 0-15-6t-6-15V64q0-9 6-15t15-6h64zm-149 0q8 0 14.5 6t6.5 15v256q0 9-6.5 15t-14.5 6h-64q-9 0-15.5-6t-6.5-15V64q0-9 6.5-15t15.5-6h64z"/>`),
		g.Group(children),
	)
}