package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAdmin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.07 5.662H15.93a4.36 4.36 0 0 0-3.776 2.18l-8.07 13.979a4.359 4.359 0 0 0 0 4.359l8.07 13.978a4.36 4.36 0 0 0 3.776 2.18h16.14a4.36 4.36 0 0 0 3.775-2.18l8.071-13.978a4.359 4.359 0 0 0 0-4.36l-8.07-13.978a4.36 4.36 0 0 0-3.776-2.18Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.32 31.785h12.185l6.093 10.552m13.484-28.89L32.989 24l6.093 10.553M7.32 16.215h12.185l6.093-10.553m2.897 10.553h-8.99L15.011 24l4.494 7.785h8.99L32.989 24l-4.494-7.785z"/>`),
		g.Group(children),
	)
}