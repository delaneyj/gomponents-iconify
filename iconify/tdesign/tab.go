package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.5 20h-21v-2h21v2Zm0-16v12h-6V4h6Zm-2 2h-2v8h2V6ZM15 4v12H9V4h6Zm-2 2h-2v8h2V6ZM7.5 4v12h-6V4h6Zm-2 2h-2v8h2V6Z"/>`),
		g.Group(children),
	)
}