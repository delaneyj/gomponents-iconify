package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lbry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.9 24.72l14.67 7.21l19.08-11.72v-1.42l-18.06-8.72L4.5 22.51v6l19.07 9.42L43.5 25.71"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.89 28l.61-2.29l-2.29-.6"/>`),
		g.Group(children),
	)
}