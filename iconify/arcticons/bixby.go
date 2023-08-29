package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bixby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5A18.5 18.5 0 1 1 5.5 24V5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.715 22.547a6.865 6.865 0 1 1-6.865 6.864v-6.864Zm-6.865 0V5.57"/>`),
		g.Group(children),
	)
}