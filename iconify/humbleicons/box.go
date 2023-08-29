package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 21l7.794-4.5v-9M12 21l-7.794-4.5v-9M12 21v-9m7.794-4.5L12 3L4.206 7.5m15.588 0L12 12M4.206 7.5L12 12"/>`),
		g.Group(children),
	)
}