package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Writer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 30.872h11.462M12 23.915h11.462M12 16.957h4.487M43.5 24.198V9.401l-12.854 7.398L43.5 24.198Zm-5.396 1.955v12.09a2.337 2.337 0 0 1-2.345 2.34H6.845a2.337 2.337 0 0 1-2.345-2.34V9.757a2.337 2.337 0 0 1 2.345-2.34h31.308"/>`),
		g.Group(children),
	)
}