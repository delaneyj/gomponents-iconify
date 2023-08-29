package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Read(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm3.465-4a4.002 4.002 0 0 0-7.339 1H2a1 1 0 1 0 0 2h1.126A4.002 4.002 0 0 0 11 12h2a4 4 0 0 0 7.874 1H22a1 1 0 1 0 0-2h-1.126a4.002 4.002 0 0 0-7.339-1h-3.07ZM15 12a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}