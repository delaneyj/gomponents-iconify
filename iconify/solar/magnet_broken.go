package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 2h2.5A1.5 1.5 0 0 1 21 3.5v2A1.5 1.5 0 0 1 19.5 7H17m0-5h-4a9.966 9.966 0 0 0-4 .832M17 2v5m0 0h-4a5 5 0 0 0-4 8m8 2h2.5a1.5 1.5 0 0 1 1.5 1.5v2a1.5 1.5 0 0 1-1.5 1.5H17m0-5h-4m4 0v5m0 0h-4C7.477 22 3 17.523 3 12c0-2.252.744-4.33 2-6.001"/>`),
		g.Group(children),
	)
}