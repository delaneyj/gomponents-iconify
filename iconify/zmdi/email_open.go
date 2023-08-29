package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m426 155l1 213q0 18-12.5 30.5T384 411H43q-18 0-30.5-12.5T0 368V155q0-24 20-37L213 5l193 113q20 13 20 37zM213 261l177-110L213 48L37 151z"/>`),
		g.Group(children),
	)
}