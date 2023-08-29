package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dharma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5h0A18.5 18.5 0 0 1 42.5 24v16.5a2 2 0 0 1-2 2h-33a2 2 0 0 1-2-2V24A18.5 18.5 0 0 1 24 5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.83 11.384h0a13.465 13.465 0 0 1 13.465 13.465V42.5h0h-26.93h0V24.849A13.465 13.465 0 0 1 23.83 11.384Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.819 28.922h0a5.29 5.29 0 0 1 5.29 5.29V42.5h0h-10.58h0v-8.288a5.29 5.29 0 0 1 5.29-5.29Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.546 20.427h0a9.636 9.636 0 0 1 9.636 9.636V42.5h0H15.91h0V30.063a9.636 9.636 0 0 1 9.636-9.636Z"/>`),
		g.Group(children),
	)
}