package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h9.11A1.89 1.89 0 0 0 16 6.11v0c0-1.615-1.894-2.486-3.12-1.435L12.5 5M3 12h14.902C19.06 12 20 12.94 20 14.098v0c0 2.152-2.853 2.91-3.92 1.041L16 15M5 16h6.11A1.89 1.89 0 0 1 13 17.89v0c0 1.615-1.894 2.486-3.12 1.435L9.5 19"/>`),
		g.Group(children),
	)
}