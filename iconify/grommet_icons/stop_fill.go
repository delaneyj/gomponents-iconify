package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4 4h16v16H4V4Zm2 2h12v12H6V6Zm2 2h8v8H8V8Zm2 2h4v4h-4v-4Zm1 1h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}