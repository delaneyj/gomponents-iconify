package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M11.5 6.5v5a3 3 0 0 1-3 3h-5a3 3 0 0 1-3-3v-5a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1Zm0 0h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2M4.5 4V2m3 2V0"/>`),
		g.Group(children),
	)
}