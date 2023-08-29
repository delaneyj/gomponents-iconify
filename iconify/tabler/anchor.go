package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v12m-8-8a8 8 0 0 0 16 0m1 0h-2M5 13H3m6-7a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/>`),
		g.Group(children),
	)
}