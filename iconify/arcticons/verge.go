package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Verge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.6 4.5l18.4 39l18.4-39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 21.41L31.98 4.5H16.02L24 21.41z"/>`),
		g.Group(children),
	)
}