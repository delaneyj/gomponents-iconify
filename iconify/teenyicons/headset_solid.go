package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.5 6a2.49 2.49 0 0 0-1.414.438a6.502 6.502 0 0 1 12.828 0A2.488 2.488 0 0 0 12.5 6A1.5 1.5 0 0 0 11 7.5v4a1.5 1.5 0 0 0 .947 1.395A1.5 1.5 0 0 1 10.5 14H8v1h2.5a2.5 2.5 0 0 0 2.458-2.042A2.5 2.5 0 0 0 15 10.5v-3a7.5 7.5 0 0 0-15 0v3A2.5 2.5 0 0 0 2.5 13A1.5 1.5 0 0 0 4 11.5v-4A1.5 1.5 0 0 0 2.5 6Z"/>`),
		g.Group(children),
	)
}