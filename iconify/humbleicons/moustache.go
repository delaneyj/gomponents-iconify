package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moustache(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12.57c-4 1.733-4-4.767-9-1.767c-5-3-5 3.5-9 1.767c1 4.233 6 4.233 9 1.233c3 3 8 3 9-1.233Z"/>`),
		g.Group(children),
	)
}