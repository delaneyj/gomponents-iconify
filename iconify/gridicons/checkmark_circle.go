package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 17.768l-4.884-4.884l1.768-1.768L11 14.232l8.658-8.658A9.98 9.98 0 0 0 12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10a9.94 9.94 0 0 0-.966-4.266L11 17.768z"/>`),
		g.Group(children),
	)
}