package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurvedConnectorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2.5 1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 0h2a3 3 0 0 1 3 3v6a3 3 0 0 0 3 3h2m0 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`),
		g.Group(children),
	)
}