package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func License(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 6h12v2H10zm0 4h12v2H10zm0 14h6v2h-6zm0-10h6v2h-6z"/><path fill="currentColor" d="M24 30H8a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2ZM8 4v24h16V4Z"/>`),
		g.Group(children),
	)
}