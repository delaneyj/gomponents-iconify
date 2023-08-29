package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPanoramaHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6.55c2.6.77 5.28 1.16 8 1.16c2.72 0 5.41-.39 8-1.16v10.91c-2.6-.77-5.28-1.16-8-1.16c-2.72 0-5.41.39-8 1.16V6.55M2 3.77v16.47s.77-.26.88-.3A26.24 26.24 0 0 1 12 18.3c3.09 0 6.18.55 9.12 1.64c.11.04.88.3.88.3V3.77s-.77.26-.88.3C18.18 5.15 15.09 5.71 12 5.71s-6.18-.56-9.12-1.64c-.11-.05-.88-.3-.88-.3z"/>`),
		g.Group(children),
	)
}