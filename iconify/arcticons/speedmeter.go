package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speedmeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.13 5.52L29.2 22.81l4.05-3.47l7.9 4.39l-8.15.62v3.13l-26.18 15L19.34 31L9.28 8.9l7 5.76l-2.16-9.14ZM29.2 22.81l-9.86 8.18m13.7-6.64l-6.28.49"/>`),
		g.Group(children),
	)
}