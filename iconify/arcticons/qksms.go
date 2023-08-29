package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qksms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 0 0 2.5 24h0A21.5 21.5 0 0 0 24 45.5h0a21.49 21.49 0 0 0 9.54-2.26a11.33 11.33 0 0 1 2.86-.79c3.61-.26 6.14 1.48 9.1 3c-1.57-3-3.31-5.49-3-9.1a11.37 11.37 0 0 1 .71-2.65h0A21.5 21.5 0 0 0 24 2.5Zm0 6.7A14.8 14.8 0 0 1 38.8 24h0a14.91 14.91 0 0 1-1.61 6.71h0a7.08 7.08 0 0 0-.49 1.81c-.18 2.49 1 4.23 2.1 6.27c-2-1.08-3.78-2.28-6.27-2.1a8 8 0 0 0-2 .54A14.76 14.76 0 0 1 9.2 24h0A14.8 14.8 0 0 1 24 9.2Z"/>`),
		g.Group(children),
	)
}