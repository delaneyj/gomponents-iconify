package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemOk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.004 2.004h-18a2.006 2.006 0 0 0-2 2v12a2.006 2.006 0 0 0 2 2h7v2h-2v2h8v-2h-2v-2h7a2.006 2.006 0 0 0 2-2v-12a2.006 2.006 0 0 0-2-2ZM10.756 13.67h-.001l-4.75-4.75L7.42 7.507l3.336 3.336l5.835-5.836l1.415 1.415Z"/>`),
		g.Group(children),
	)
}