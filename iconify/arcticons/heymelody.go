package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heymelody(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 15.643v16.714A11.143 11.143 0 0 0 35.143 43.5h0V15.643A11.143 11.143 0 1 0 24 26.786"/>`),
		g.Group(children),
	)
}