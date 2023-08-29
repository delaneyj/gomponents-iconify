package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forkclient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.22 17.78l7.06-7.06M25 12.56l7.07-7.06M35.44 23l7.06-7.07M35.44 23A7.38 7.38 0 0 1 25 12.56M5.5 42.5L25 23"/>`),
		g.Group(children),
	)
}