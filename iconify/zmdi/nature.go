package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 304v83h128v42H0v-42h128v-84q-53-9-88.5-50.5T4 156Q4 94 47.5 50T153 6t105.5 44T302 156q0 57-37.5 99T171 304z"/>`),
		g.Group(children),
	)
}