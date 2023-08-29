package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m141 132l-31 30L0 51l30-30zm83-111h117v118l-43-44L30 363L0 333L268 65zm7 201l67 67l43-44v118H224l44-44l-67-67z"/>`),
		g.Group(children),
	)
}