package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 9l-8 6h2v5m6-11l8 6h-2v5M12 9V4m6 16h3m-3 0h-4M3 20h3m0 0h4m0-14h4m-4 14v-3a2 2 0 0 1 2-2v0a2 2 0 0 1 2 2v3m-4 0h4"/>`),
		g.Group(children),
	)
}