package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloatingIp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 11a5.008 5.008 0 0 0-4.899 4H11.9a5 5 0 1 0 0 2h8.2a5 5 0 1 0 4.9-6Zm0 8a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Z"/>`),
		g.Group(children),
	)
}