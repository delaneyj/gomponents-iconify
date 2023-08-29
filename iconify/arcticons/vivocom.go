package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vivocom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.517 28.614a3.493 3.493 0 0 1-3.482-3.482v-2.264a3.483 3.483 0 1 1 6.965 0v2.264a3.493 3.493 0 0 1-3.483 3.482ZM8 19.386l3.483 9.229l3.482-9.23m7.307.001l3.483 9.229l3.482-9.23m-10.618.001v9.229"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}