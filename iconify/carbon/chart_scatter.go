package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartScatter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 30H4a2 2 0 0 1-2-2V2h2v26h26Z"/><circle cx="10" cy="22" r="2" fill="currentColor"/><circle cx="14" cy="15" r="2" fill="currentColor"/><circle cx="22" cy="15" r="2" fill="currentColor"/><circle cx="26" cy="6" r="2" fill="currentColor"/><circle cx="14" cy="8" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}