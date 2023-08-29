package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thelife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 41.875l17.469-9.75v-16.25h-6.5v-4.063h6.5V6.125h4.062v5.688h6.5v4.062h-6.5v16.25l17.469 9.75"/>`),
		g.Group(children),
	)
}