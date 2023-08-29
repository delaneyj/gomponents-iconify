package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TiktokOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 0v11A3.5 3.5 0 1 1 6 7.5m8-2A4.5 4.5 0 0 1 9.5 1"/>`),
		g.Group(children),
	)
}