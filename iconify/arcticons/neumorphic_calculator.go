package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeumorphicCalculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5v37M5.5 24h37m-13.773-8.574h9M30.29 28.692l6 9m-25.034-8.37l8 7.937m-.039-8l-7.961 7.969m4-26.302v9m-4.5-4.5h9"/>`),
		g.Group(children),
	)
}