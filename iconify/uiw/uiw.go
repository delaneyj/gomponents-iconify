package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uiw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m5.976 0l4.475 3.418l-1.71 5.531H3.21L1.5 3.42L5.976 0Zm0 20L1.5 16.582l1.71-5.531h5.532l1.709 5.53L5.976 20ZM18.5 12.968l-5.261 1.797l-3.252-4.705l3.252-4.705L18.5 7.152v5.816Z"/>`),
		g.Group(children),
	)
}