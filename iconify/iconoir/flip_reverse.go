package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipReverse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 12v2m-3.75-2l-.938 2m2.813 6H9.5v-2m-5.625 2H2l.938-2M7.625 8L9.5 4v4m5 12H22L14.5 4v16Z"/>`),
		g.Group(children),
	)
}