package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cookie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm-2 4a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm5.5 1a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3zM11 13a2 2 0 0 0 0 4a2 2 0 0 0 0-4zm6 2a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm5 1a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm-9.5 3a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3zm7 1a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3z"/>`),
		g.Group(children),
	)
}