package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yourhour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 15.637V30.5a13 13 0 0 1-26 0v-8.363"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 23.064a8.359 8.359 0 1 1-8.355 8.363A8.355 8.355 0 0 1 24 23.064ZM17.5 8.21v13.927M24 4.5v14.855M30.5 8.21v13.927"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 25.855v6.5h4.645"/>`),
		g.Group(children),
	)
}