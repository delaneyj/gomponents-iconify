package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Npm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 0h24v24H0z"/><path d="M16.718 7.928h-4.513V20.25H4V3h16v17.249h-3.282V7.93Z"/></g>`),
		g.Group(children),
	)
}