package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTwitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20.5 4.5l-2 1.5M21 7h-2M3.5 18c11.5 4.5 17-7 15-11s-7-1.5-5.5 3c-3.5 0-7-1-9-3.5c-1 3.5.5 8 4 9.5c-1.065 1.352-1.795 1.703-4.5 2Z"/>`),
		g.Group(children),
	)
}