package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 4l-5 6.667M6 20l5-6.667m2-2.666L8.3 4.4a1 1 0 0 0-.8-.4H4l7 9.333m2-2.666L20 20h-3.5a1 1 0 0 1-.8-.4L11 13.333"/>`),
		g.Group(children),
	)
}