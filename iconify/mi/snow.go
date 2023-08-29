package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2a6 6 0 0 0-5.986 6.41a5 5 0 0 0-1.322 8.34a1 1 0 1 0 1.324-1.5a3.002 3.002 0 0 1 1.324-5.178a1 1 0 0 0 .757-1.193A4 4 0 1 1 14.92 7.2a1 1 0 0 0 .999.8H16a4 4 0 0 1 2.4 7.2a1 1 0 0 0 1.201 1.6a6 6 0 0 0-2.93-10.762A6.002 6.002 0 0 0 11 2zm3.5 15a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm-3.5-.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0zm4 3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0zm-5 1a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z"/>`),
		g.Group(children),
	)
}