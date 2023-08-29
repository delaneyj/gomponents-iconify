package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 17h-2a4 4 0 0 0-4 4v0a4 4 0 0 0 4 4h2m10 8h-3a4 4 0 0 0-4 4v0a4 4 0 0 0 4 4h3m5-34h-3a4 4 0 0 0-4 4v0a4 4 0 0 0 4 4h3m-3 0h16M12 25h30m-20 8h12"/>`),
		g.Group(children),
	)
}