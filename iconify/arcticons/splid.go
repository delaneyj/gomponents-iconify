package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Splid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.077 6.513A11.594 11.594 0 1 0 21.02 25.677m-7.097 15.81A11.594 11.594 0 1 0 26.98 22.323"/>`),
		g.Group(children),
	)
}