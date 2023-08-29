package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneHundredEighty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m9.5 18.917l3.492-1.902m0 0v13.97m16.253-4.628a4.627 4.627 0 1 0 9.255 0v-4.714a4.627 4.627 0 1 0-9.255 0v4.714ZM20.106 24a3.492 3.492 0 0 0-3.492 3.492h0a3.492 3.492 0 0 0 3.492 3.493h2.27a3.492 3.492 0 0 0 3.493-3.493h0A3.492 3.492 0 0 0 22.376 24"/><path d="M22.376 24a3.492 3.492 0 0 0 3.493-3.492h0a3.492 3.492 0 0 0-3.493-3.493h-2.27a3.492 3.492 0 0 0-3.492 3.493h0A3.492 3.492 0 0 0 20.106 24m0 0h2.27"/></g>`),
		g.Group(children),
	)
}