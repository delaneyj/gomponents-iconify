package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opensync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.711 29.711a8.077 8.077 0 1 0 0-11.422L18.29 29.71a8.077 8.077 0 1 1 0-11.422"/>`),
		g.Group(children),
	)
}