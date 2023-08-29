package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageInPictureLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 15v-2a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h2m5-5v3a2 2 0 0 1-2 2H8m5-5c-2 0-5 1-5 5m-1-6h.01M4 8h.001M20 8h.001M20 12h.001M20 16h.001M20 20h.001M16 20h.001M4 4h.001M8 4h.001M12 4h.001M16 4h.001M20 4h.001"/>`),
		g.Group(children),
	)
}