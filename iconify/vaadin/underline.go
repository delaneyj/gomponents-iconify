package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 15h12v1H2v-1zm9-15v8.4C11 9.9 9.9 11 8.4 11h-.8C6.1 11 5 9.9 5 8.4V0H2v8.4C2 11.5 4.5 14 7.6 14h.9c3.1 0 5.6-2.5 5.6-5.6V0H11z"/>`),
		g.Group(children),
	)
}