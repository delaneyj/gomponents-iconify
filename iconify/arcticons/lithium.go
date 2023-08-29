package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lithium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 35.5v-23a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v23a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2ZM24 10.5v27"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40 10.544v12l-2.5-2l-2.5 2v-12"/>`),
		g.Group(children),
	)
}