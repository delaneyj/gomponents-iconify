package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.423 5.5v37H19.86V30.845h8.281V42.5h12.437V21.487l-24.904-7.122V5.5h-8.25Z"/>`),
		g.Group(children),
	)
}