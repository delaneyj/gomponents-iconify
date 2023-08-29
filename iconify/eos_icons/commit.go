package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 11h-6.141a3.981 3.981 0 0 0-7.718 0H2v2h6.141a3.981 3.981 0 0 0 7.718 0H22Z"/>`),
		g.Group(children),
	)
}