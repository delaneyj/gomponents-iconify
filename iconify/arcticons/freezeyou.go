package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freezeyou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 38.024c1.032-3.529 3.47-8.057 5.986-8.057c4.772 0 6.409 10.55 13.674 10.55c7.844 0 16.475-15.047 17.34-33.034"/>`),
		g.Group(children),
	)
}