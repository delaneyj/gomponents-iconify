package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sidebar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M21 20H7V4h14v16Zm-2-2H9V6h10v12Z" clip-rule="evenodd"/><path d="M3 20h2V4H3v16Z"/></g>`),
		g.Group(children),
	)
}