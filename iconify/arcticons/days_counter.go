package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DaysCounter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.477 16.102l-.046 10.74l9.677 8.239m-8.511 10.294H45.5V28.372"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.596 44.98A21.06 21.06 0 0 1 2.599 22.259A21.581 21.581 0 0 1 45.5 21.48"/>`),
		g.Group(children),
	)
}