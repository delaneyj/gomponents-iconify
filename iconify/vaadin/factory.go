package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Factory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.4 1.3c-.6.3-.8 1.1-.4 1.5c.5-.9 1.3-.6 2.5.4c.8.7 1.9.1 1.9.1s.2 1.2 1.7 1.4c1.7.2 2.3-.8 2.3-.8s.4 1 1.9.4c1.1-.4.7-1.1.7-1.1s1 0 1-.7c0-.9-1.1-.8-1.1-.8s.2-1-.9-1.1c-1-.1-1.3.5-1.3.5S12.4 0 10.9 0C9.5 0 9 1.3 9 1.3S8.6.7 7.4.7c-.9 0-1.3.7-1.3.7S5 .9 4.4 1.3z"/><path fill="currentColor" d="M12 12.1V10l-4 2.1V10H5.6L5 3H3l-.6 7H0v6h16v-6l-4 2.1zM6 14H2v-2h4v2z"/>`),
		g.Group(children),
	)
}