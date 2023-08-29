package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.62 24.332a2.88 2.88 0 0 1 2.88-2.88h0m-2.881 0v7.633"/><circle cx="11.373" cy="18.923" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.696 19.075v8.57c0 .796.645 1.44 1.44 1.44h.433m-1.873-7.633h1.873m-15.069 0h1.873v7.633m8.561 0v-4.753a2.88 2.88 0 0 0-2.88-2.88h0a2.88 2.88 0 0 0-2.881 2.88v4.753m-.001-4.753v-2.88m18.297 6.18a2.88 2.88 0 0 1-2.503 1.453h0a2.88 2.88 0 0 1-2.881-2.88v-1.873a2.88 2.88 0 0 1 2.88-2.88h0a2.88 2.88 0 0 1 2.881 2.88v.936h-5.761"/>`),
		g.Group(children),
	)
}