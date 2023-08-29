package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Floccus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.783 35.573h23.98c13.015 0 10.837-19.586-2.216-15.243c0-10.876-19.586-10.876-19.586 2.178C2.073 20.33 2.073 35.574 10.783 35.574Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.26 26.434l4.648 4.596l7.478-9.564"/>`),
		g.Group(children),
	)
}