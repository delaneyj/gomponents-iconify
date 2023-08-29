package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindyRain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 5a1 1 0 1 0 0 2h5v2h-5a3 3 0 1 1 0-6h4v2h-4ZM5.5 8a1.5 1.5 0 1 0 0 3H12v2H5.5a3.5 3.5 0 1 1 0-7H8v2H5.5ZM20 6.996h2.004V9H20V6.996Zm0 4h2.004V13H20v-2.004ZM14 11h4v2h-4v-2Zm2 3.996h2.004V17H16v-2.004ZM5 18a3 3 0 0 1 3-3h1v2H8a1 1 0 1 0 0 2h4.5v2H8a3 3 0 0 1-3-3Zm6-3h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}