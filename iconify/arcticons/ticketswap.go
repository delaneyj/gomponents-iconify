package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticketswap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.12 33.11a12.88 12.88 0 0 1-18.21 0h0L4.47 22.71L9 18.14a2.77 2.77 0 0 1 3.9 0l8.5 8.46a3.67 3.67 0 0 0 6.2-1.6l6.78 6.77a13.86 13.86 0 0 1-1.26 1.34ZM39 29.86a2.77 2.77 0 0 1-3.9 0l-8.49-8.46a3.67 3.67 0 0 0-6.2 1.65l-6.78-6.76a14.53 14.53 0 0 1 1.2-1.4a12.88 12.88 0 0 1 18.21 0h0l10.41 10.4Z"/>`),
		g.Group(children),
	)
}