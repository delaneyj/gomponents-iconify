package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m.29 12.71l1.42-1.42l2.22 2.22l8.3-10.14l1.54 1.26l-9.7 11.86zM12 10h5v2h-5zm-3 4h5v2H9zm6-8h5v2h-5z"/>`),
		g.Group(children),
	)
}