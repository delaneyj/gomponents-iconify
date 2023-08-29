package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiNoneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 11.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`),
		g.Group(children),
	)
}