package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M112 152h288v32H112zm0 88h288v32H112zm0 88h152v32H112z"/><path fill="currentColor" d="M480 48H32v416h448Zm-32 384H64V80h384Z"/>`),
		g.Group(children),
	)
}