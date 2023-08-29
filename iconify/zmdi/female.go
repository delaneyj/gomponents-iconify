package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Female(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 429H64V301H0l54-162q4-14 15.5-22t25.5-8h2q14 0 25 8t16 22l54 162h-64v128zM96 88q-18 0-30.5-12.5T53 45.5t12.5-30T96 3t30.5 12.5t12.5 30t-12.5 30T96 88z"/>`),
		g.Group(children),
	)
}