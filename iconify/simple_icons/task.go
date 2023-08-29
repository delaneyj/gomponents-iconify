package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Task(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.857 18.013L11.736 24V12.456L1.857 6.468Zm20.286 0V6.468l-9.879 5.988V24Zm-.246-12.014L12 0L2.103 5.999L12 11.998Z"/>`),
		g.Group(children),
	)
}