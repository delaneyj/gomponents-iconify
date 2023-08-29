package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustVerticalAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3 0v8.05a2.5 2.5 0 0 0 0 4.9V15h1v-2.05a2.5 2.5 0 0 0 0-4.9V0H3Zm8 0v2.05a2.5 2.5 0 0 0 0 4.9V15h1V6.95a2.5 2.5 0 0 0 0-4.9V0h-1Z"/>`),
		g.Group(children),
	)
}