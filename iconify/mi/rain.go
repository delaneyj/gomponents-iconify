package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2a6 6 0 0 0-5.986 6.41a5 5 0 0 0-1.322 8.34a1 1 0 1 0 1.324-1.5a3.002 3.002 0 0 1 1.324-5.178a1 1 0 0 0 .757-1.193A4 4 0 1 1 14.92 7.2a1 1 0 0 0 .999.8H16a4 4 0 0 1 2.4 7.2a1 1 0 0 0 1.201 1.6a6 6 0 0 0-2.93-10.762A6.002 6.002 0 0 0 11 2zm1.949 13.316a1 1 0 0 0-1.898-.632l-2 6a1 1 0 0 0 1.898.632l2-6zm3.367-2.265a1 1 0 0 1 .633 1.265l-2 6a1 1 0 0 1-1.898-.632l2-6a1 1 0 0 1 1.265-.633zM9.45 14.316a1 1 0 0 0-1.898-.632l-2 6a1 1 0 0 0 1.898.632l2-6z"/>`),
		g.Group(children),
	)
}