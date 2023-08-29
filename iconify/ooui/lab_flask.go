package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabFlask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 7.61V3h1V1H6v2h1v4.61l-5.86 9.88A1 1 0 0 0 2 19h16a1 1 0 0 0 .86-1.51zm-4.2.88a1 1 0 0 0 .2-.6V3h2v4.89a1 1 0 0 0 .14.51l2.14 3.6H6.72z"/>`),
		g.Group(children),
	)
}