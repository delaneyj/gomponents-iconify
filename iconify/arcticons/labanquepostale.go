package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Labanquepostale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 31.06l39-9.3H19.741s-.897 4.145-15.241 9.3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.213 24.45a48.741 48.741 0 0 1-5.846 6.61H4.5m15.24-9.302c-1.121-1.68-4.483-3.698-8.181-4.818h22.526a2.35 2.35 0 0 1 .772 1.96a6.614 6.614 0 0 1-.93 2.858"/>`),
		g.Group(children),
	)
}