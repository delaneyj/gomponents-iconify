package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0ZM9.928 8.124l4.419 7.015l2.581-3.975L23.966 22H1.188l8.74-13.876Zm0 3.752L4.812 20h15.47l-3.354-5.164l-2.607 4.014l-4.393-6.974Z"/>`),
		g.Group(children),
	)
}