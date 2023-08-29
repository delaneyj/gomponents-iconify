package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jisho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.688 15.842c4.333-.466 9.369-1.717 14.125-4.197"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.734 5.5c-8.135 20.897-.572 30.312 1.638 34.312"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.3 23.89c-6.879 18.127-10.062 18.508-13.77 15.367c-2.34-1.983-.595-5.58 2.86-8.575c4.215-3.657 11.845-4.963 15.304-3.942c11.595 3.423 2.885 12.904-.203 15.76"/>`),
		g.Group(children),
	)
}