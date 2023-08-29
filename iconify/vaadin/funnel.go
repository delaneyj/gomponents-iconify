package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Funnel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 11h4v4H6v-4zm7.6-6L16 1H0l2.4 4h11.2zM3 6l2.4 4h5.2L13 6H3z"/>`),
		g.Group(children),
	)
}