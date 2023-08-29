package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OhNo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M22.5 14A8.5 8.5 0 1 1 14 5.5a8.5 8.5 0 0 1 8.5 8.5Zm20 0A8.5 8.5 0 1 1 34 5.5a8.5 8.5 0 0 1 8.5 8.5Zm0 20a8.5 8.5 0 1 1-8.5-8.5a8.5 8.5 0 0 1 8.5 8.5Z"/>`),
		g.Group(children),
	)
}