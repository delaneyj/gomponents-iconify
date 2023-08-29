package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobephotoshopmix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40 5.5H7a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="21" cy="21" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.5" cy="26.5" r="7.5" fill="none" stroke="currentColor" stroke-dasharray="1 1.78" stroke-dashoffset="1.6" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}