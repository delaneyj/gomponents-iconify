package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WordCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 16h12v2H10zm0 4h8v2h-8z"/><path fill="currentColor" d="M16 7a8.023 8.023 0 0 1 7.865 6.494l.259 1.346l1.349.244A5.502 5.502 0 0 1 24.508 26H7.495a5.502 5.502 0 0 1-.97-10.916l1.35-.244l.259-1.346A8.026 8.026 0 0 1 16 7m0-2a10.024 10.024 0 0 0-9.83 8.116A7.502 7.502 0 0 0 7.491 28h17.017a7.502 7.502 0 0 0 1.32-14.884A10.023 10.023 0 0 0 15.989 5Z"/>`),
		g.Group(children),
	)
}