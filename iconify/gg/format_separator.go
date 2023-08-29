package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatSeparator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16 5a1 1 0 1 0 0-2H8a1 1 0 1 0 0 2h8Zm0 2a1 1 0 1 1 0 2H8a1 1 0 1 1 0-2h8Zm1 5a1 1 0 0 1-1 1H8a1 1 0 1 1 0-2h8a1 1 0 0 1 1 1Zm-1 9a1 1 0 1 0 0-2H8a1 1 0 1 0 0 2h8Z" opacity=".5"/><path fill-rule="evenodd" d="M21 16a1 1 0 0 1-1 1H4a1 1 0 1 1 0-2h16a1 1 0 0 1 1 1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}