package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 4.5l7-4l7 4v.72a9.651 9.651 0 0 1-7 9.28a9.651 9.651 0 0 1-7-9.28V4.5Z"/>`),
		g.Group(children),
	)
}