package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.5 0h-7C3.675 0 3 .675 3 1.5v13c0 .825.675 1.5 1.5 1.5h7c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5zM6 .75h4v.5H6v-.5zM8 15a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm4-3H4V2h8v10z"/>`),
		g.Group(children),
	)
}