package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConvertToCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.5 12H19a4 4 0 0 1 0-8h.084A4.786 4.786 0 0 1 23 2a4.982 4.982 0 0 1 4.654 3.203A3.467 3.467 0 0 1 30 8.5a3.504 3.504 0 0 1-3.5 3.5zM19 6a2 2 0 0 0 0 4h7.5A1.502 1.502 0 0 0 28 8.5a1.486 1.486 0 0 0-1.278-1.474l-.661-.099l-.162-.649a2.957 2.957 0 0 0-5.487-.712L20.114 6zm4 8l-5 5l1.414 1.414L22 17.832V24h-6v2h6a2.002 2.002 0 0 0 2-2v-6.168l2.587 2.583L28 19zM2 20h10v2H2zm0 4h10v2H2zm0 4h10v2H2z"/>`),
		g.Group(children),
	)
}