package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneTextIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.99 19h2.42l1.27-3.58h5.65L12.59 19h2.42L9.75 5h-2.5L1.99 19zm4.42-5.61L8.44 7.6h.12l2.03 5.79H6.41zM20 11h3v2h-3v3h-2v-3h-3v-2h3V8h2v3z"/>`),
		g.Group(children),
	)
}