package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.637 19.561l-2.219 8.878l-2.22-8.878l-2.219 8.878l-2.22-8.878m-5.881 8.878l-2.22-8.878l-2.219 8.878l-2.22-8.878L8 28.439m14.759-8.878L19.819 24l-2.941-4.439m2.941 8.878V24M40 19.561L37.059 24l-2.941-4.439m2.941 8.878V24m-3.502 1.469H30.16m1.477-5.908l2.885 8.878"/>`),
		g.Group(children),
	)
}