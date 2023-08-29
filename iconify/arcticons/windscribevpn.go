package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windscribevpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 44.5l-14.5-6l-6-14.5l6-14.5l14.5-6l14.5 6l6 14.5l-6 14.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.82 14.48V33.7h3.07a3.43 3.43 0 0 0 2.2-1h0L24 27.34l5.9 5.38h0a3.52 3.52 0 0 0 2.21 1h3.07V14.48"/>`),
		g.Group(children),
	)
}