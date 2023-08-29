package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nexttracks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a13.095 13.095 0 0 0-13.095 13.095c0 10.248 10.024 22.6 12.62 25.628a.791.791 0 0 0 1.208-.006c2.551-3.037 12.362-15.38 12.362-25.622A13.095 13.095 0 0 0 24 4.5Zm0 22.595a9.5 9.5 0 1 1 9.5-9.5a9.5 9.5 0 0 1-9.5 9.5Z"/>`),
		g.Group(children),
	)
}