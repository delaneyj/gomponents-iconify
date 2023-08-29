package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retrica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM24 14.84A9.16 9.16 0 0 1 33.16 24h0A9.16 9.16 0 0 1 24 33.16h0A9.16 9.16 0 0 1 14.84 24h0A9.16 9.16 0 0 1 24 14.84Zm15.48 24.64h-6.32v-6.32h6.32Zm-3.16 0v-6.32m-3.16 3.16h6.32"/>`),
		g.Group(children),
	)
}