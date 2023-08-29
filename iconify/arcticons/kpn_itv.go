package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KpnItv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.115 22.06l-11.704-6.757c-1.493-.861-3.359.216-3.359 1.94v13.515c0 1.723 1.866 2.8 3.359 1.939l11.704-6.758c1.493-.862 1.493-3.016 0-3.878Z"/>`),
		g.Group(children),
	)
}