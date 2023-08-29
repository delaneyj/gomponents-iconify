package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plutus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2.25 2.25 0 0 0-2 2v33a2.25 2.25 0 0 0 2 2h33a2.25 2.25 0 0 0 2-2v-33a2.25 2.25 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8 40l12-13l8 9l12-20"/>`),
		g.Group(children),
	)
}