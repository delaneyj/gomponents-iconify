package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApiRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 14l-2-2l2-2l2 2l-2 2ZM9.875 8.125l-2.5-2.5l3.2-3.2q.3-.3.675-.45t.75-.15q.375 0 .75.15t.675.45l3.2 3.2l-2.5 2.5L12 6L9.875 8.125Zm-4.25 8.5l-3.2-3.2q-.3-.3-.45-.675t-.15-.75q0-.375.15-.75t.45-.675l3.2-3.2l2.5 2.5L6 12l2.125 2.125l-2.5 2.5Zm12.75 0l-2.5-2.5L18 12l-2.125-2.125l2.5-2.5l3.2 3.2q.3.3.45.675t.15.75q0 .375-.15.75t-.45.675l-3.2 3.2Zm-7.8 4.95l-3.2-3.2l2.5-2.5L12 18l2.125-2.125l2.5 2.5l-3.2 3.2q-.3.3-.675.45t-.75.15q-.375 0-.75-.15t-.675-.45Z"/>`),
		g.Group(children),
	)
}