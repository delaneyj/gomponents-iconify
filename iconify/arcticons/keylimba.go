package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keylimba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="16.423" cy="20.153" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.423" ry="2.692"/><ellipse cx="31.577" cy="20.153" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.423" ry="2.692"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.671 27.46c-1.328 1.84-3.335 3.08-6.671 3.08s-5.343-1.238-6.671-3.08"/>`),
		g.Group(children),
	)
}