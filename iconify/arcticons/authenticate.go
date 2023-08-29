package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Authenticate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.17 3.42L15.49 18.83a4.3 4.3 0 0 0-1.58 5.87h0a4.3 4.3 0 0 0 5.88 1.58l7.9-4.56a4.29 4.29 0 0 1 5.87 1.58h0A4.29 4.29 0 0 1 32 29.17L5.31 44.58"/>`),
		g.Group(children),
	)
}