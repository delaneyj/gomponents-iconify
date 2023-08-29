package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FptIviec(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.024 13.517l4.687-7.092l4.687 7.092l-4.687 7.091l-4.687-7.091Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.384 5.5l3.7 8.387h13.32L24.018 29.55l-6.784-11.1l-4.933 7.092L21.55 42.5h4.933l16.342-28.613L38.818 5.5H15.384Z"/>`),
		g.Group(children),
	)
}