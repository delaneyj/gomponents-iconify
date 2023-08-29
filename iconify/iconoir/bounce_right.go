package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BounceRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM4 15.5c3-1 5.5-.5 8 4.5c.5-3 2-7.5 3.5-10"/>`),
		g.Group(children),
	)
}