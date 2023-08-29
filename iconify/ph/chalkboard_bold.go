package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChalkboardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 188h-4V56a20 20 0 0 0-20-20H40a20 20 0 0 0-20 20v132h-4a12 12 0 0 0 0 24h224a12 12 0 0 0 0-24ZM44 60h168v128h-16v-28a12 12 0 0 0-12-12h-64a12 12 0 0 0-12 12v28H84v-88h88v20a12 12 0 0 0 24 0V88a12 12 0 0 0-12-12H72a12 12 0 0 0-12 12v100H44Zm128 128h-40v-16h40Z"/>`),
		g.Group(children),
	)
}