package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Task(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<rect width="17.563" height="14.478" x="1.23" y="1035.052" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.034" rx="2.759" transform="translate(55.328 -99820.702) scale(96.7529)"/>`),
		g.Group(children),
	)
}