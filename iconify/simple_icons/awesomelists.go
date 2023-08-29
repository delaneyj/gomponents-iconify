package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Awesomelists(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m24 11.438l-6.154-5.645l-.865.944l5.128 4.7H1.895l5.128-4.705l-.865-.943l-6.154 5.649H0v3.72c0 1.683 1.62 3.053 3.61 3.053h3.795c1.99 0 3.61-1.37 3.61-3.051v-2.446h1.97v2.446c0 1.68 1.62 3.051 3.61 3.051h3.794c1.99 0 3.61-1.37 3.61-3.051v-3.721z"/>`),
		g.Group(children),
	)
}