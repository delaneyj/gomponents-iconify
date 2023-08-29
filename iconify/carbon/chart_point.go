package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 30H4a2.002 2.002 0 0 1-2-2V2h2v26h26Z"/><circle cx="9" cy="6" r="3" fill="currentColor"/><circle cx="9" cy="22" r="3" fill="currentColor"/><circle cx="18" cy="14" r="3" fill="currentColor"/><path fill="currentColor" d="M9 17a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1Z"/><circle cx="27" cy="6" r="3" fill="currentColor"/><circle cx="27" cy="22" r="3" fill="currentColor"/><path fill="currentColor" d="M27 17a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1zm-9 13a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1zm0-14a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}