package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tinytinyrss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.1 10.5l-5 22.86a3.23 3.23 0 0 0 3.24 4.14a3.58 3.58 0 0 0 1.75-.28m-5.44-21.68h8.7"/>`),
		g.Group(children),
	)
}