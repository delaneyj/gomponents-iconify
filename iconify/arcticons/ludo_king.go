package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LudoKing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7.875" cy="30.887" r="5.375"/><circle cx="18.625" cy="30.887" r="5.375"/><circle cx="29.375" cy="30.887" r="5.375"/><circle cx="40.125" cy="30.887" r="5.375"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.5 11.738l-3.75 5l-3.75-5l-3.75 5l-3.75-5v10.299h15V11.738z"/>`),
		g.Group(children),
	)
}