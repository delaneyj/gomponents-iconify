package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telegram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1189 1307l147-693q9-44-10.5-63t-51.5-7L410 877q-29 11-39.5 25t-2.5 26.5t32 19.5l221 69l513-323q21-14 32-6q7 5-4 15l-415 375l-16 228q23 0 45-22l108-104l224 165q64 36 81-38zm603-411q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348z"/>`),
		g.Group(children),
	)
}