package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noorulhuda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c-4.17-4.036-11-2.763-11-7.15v-24.7c0-4.387 6.83-3.114 11-7.15c4.17 4.036 11 2.763 11 7.15v24.7c0 4.387-6.83 3.114-11 7.15Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 35h9V13h-9m9 0h8m-8 22h8m9-22h-9v22h9"/>`),
		g.Group(children),
	)
}