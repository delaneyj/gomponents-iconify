package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4a2 2 0 0 1 2-2h4v2h-4v2h9v14H2V6h9V4ZM4 8v10h16V8H4Zm5 2v2h2v2H9v2H7v-2H5v-2h2v-2h2Zm6 0h2.003v2h2v2.004h-2v2H15v-2h-2V12h2v-2Zm.004 2.004V14H17v-1.996h-1.996Z"/>`),
		g.Group(children),
	)
}