package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M14.499 3.996a2.2 2.2 0 0 1 1.556.645l3.302 3.301a2.2 2.2 0 0 1 0 3.113l-.567.567l.043.192a8.5 8.5 0 0 1-3.732 8.83l-.23.144a8.5 8.5 0 1 1-2.687-15.623l.192.042l.567-.566a2.2 2.2 0 0 1 1.362-.636zM10 9a4 4 0 0 0-4 4a1 1 0 0 0 2 0a2 2 0 0 1 2-2a1 1 0 0 0 0-2z"/><path fill="currentColor" d="M21 2a1 1 0 0 1 .117 1.993L21 4h-1c0 .83-.302 1.629-.846 2.25L19 6.413l-1.293 1.293a1 1 0 0 1-1.497-1.32l.083-.094L17.586 5c.232-.232.375-.537.407-.86L18 4a2 2 0 0 1 1.85-1.995L20 2h1z"/></g>`),
		g.Group(children),
	)
}