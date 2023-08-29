package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tripadvisor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="14.25" cy="25.938" r="9.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.75" cy="25.938" r="9.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.25" cy="25.938" r="3.079" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.75" cy="25.938" r="3.079" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.729 16.188H4.5l2.858 2.858"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.642 19.046C35.455 13.858 30.096 12.313 24 12.313s-11.455 1.546-16.642 6.732m19.5 13.785L24 35.688l-2.858-2.858"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.271 16.188H43.5l-2.858 2.858"/>`),
		g.Group(children),
	)
}