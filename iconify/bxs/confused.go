package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Confused(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm-5 8.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 7 10.5zm1.124 6.492l-.248-1.984l8-1l.248 1.984l-8 1zm7.369-5.006a1.494 1.494 0 1 1 .001-2.987a1.494 1.494 0 0 1-.001 2.987z"/>`),
		g.Group(children),
	)
}