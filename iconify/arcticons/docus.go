package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Docus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 15v26.55a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95H18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18 33V15h4.06A7.88 7.88 0 0 1 30 22.87v2.25A7.88 7.88 0 0 1 22.09 33Z"/>`),
		g.Group(children),
	)
}