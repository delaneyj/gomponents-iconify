package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eboox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 12.268v23.243h14.511L24 38.974l4.989-3.463H43.5V12.268"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.765 9.026L24 11.331v20.521L8.765 29.547V9.026zm30.469 0L24 11.331v20.521l15.234-2.305V9.026z"/>`),
		g.Group(children),
	)
}