package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChimneyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.28 2H1.72l1.27 5.088L1.926 22h20.073V10h-9.66l-.326-2.932L13.28 2Zm-3.175 6l.222 2H8v10H4.074L4.93 8h5.173ZM4.78 6l-.5-2h6.439l-.5 2H4.78ZM10 20v-8h10v8h-4v-4h-2v4h-4Z"/>`),
		g.Group(children),
	)
}