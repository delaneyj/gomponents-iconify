package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MastodonRedirect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.048 26.787v14.627h-8.042V29.809a5.52 5.52 0 0 0-1.255-3.495l-7.03-8.528l-6.221 6.22V6.614h17.393L17.806 11.7l6.06 7.341l.512.62l1.97 2.39a7.488 7.488 0 0 1 1.7 4.736Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.865 19.041L36.32 6.586l6.18 6.18l-14.452 14.466"/>`),
		g.Group(children),
	)
}