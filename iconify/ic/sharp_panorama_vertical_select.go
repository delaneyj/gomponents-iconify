package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPanoramaVerticalSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.49 11.99c0-5.25 1.54-9.01 1.92-10H3.59c.76 2.16 1.9 5.21 1.9 10c0 4.78-1.17 7.91-1.9 10H20.4c-.74-2.08-1.91-5.23-1.91-10z"/>`),
		g.Group(children),
	)
}