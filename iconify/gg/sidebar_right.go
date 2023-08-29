package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SidebarRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 4h14v16H3V4Zm2 2h10v12H5V6Z" clip-rule="evenodd"/><path d="M21 4h-2v16h2V4Z"/></g>`),
		g.Group(children),
	)
}