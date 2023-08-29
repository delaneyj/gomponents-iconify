package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchFont(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 19.5L21 21m-7-4a3 3 0 1 0 6 0a3 3 0 0 0-6 0ZM9 5v12m0 0H7m2 0h2m4-10V5H3v2"/>`),
		g.Group(children),
	)
}