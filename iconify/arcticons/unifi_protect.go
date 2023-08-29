package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnifiProtect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="13.123" height="14.966" x="17.439" y="7.067" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.561"/><rect width="17.624" height="20.1" x="15.188" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8.812"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.747 15.788v24.796a2.406 2.406 0 0 1-2.023 2.376h0a42.51 42.51 0 0 1-13.512 0h0a2.406 2.406 0 0 1-2.024-2.375V15.788"/>`),
		g.Group(children),
	)
}