package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Biblio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 15.73v23.72m6.5-26.9v26.9M17.5 19v20.45M24 8.55v30.9M30.5 15v24.45m6.5-17.9v17.9M43.5 11v28.45"/>`),
		g.Group(children),
	)
}