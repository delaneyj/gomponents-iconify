package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Columns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3v18c0 .6.4 1 1 1h8V2H3c-.6 0-1 .4-1 1zm19-1h-8v20h8c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}