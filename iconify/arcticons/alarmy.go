package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarmy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.945" cy="24.129" r="18.411" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".869" d="M7.656 15.009a5.754 5.754 0 1 1 8.598-7.647m15.938.302a5.754 5.754 0 1 1 8.158 8.002"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.385 13.1v11.873l6.638 5.648"/>`),
		g.Group(children),
	)
}