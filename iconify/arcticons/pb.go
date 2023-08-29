package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.88 5.86h6.62v5.89h-4.78l-7.33 7.34l7 7.05a1.43 1.43 0 0 1 0 2l-9.25 9.25a1.43 1.43 0 0 1-1.86.13l-2.06 2.07a8.69 8.69 0 0 1-12.21 0l-.24-.24a1.43 1.43 0 0 1-1.94-.06l-4-4a1.45 1.45 0 0 1 0-2L14 25.27l-.09-.09a1.44 1.44 0 0 1 0-2l9.25-9.26a1.45 1.45 0 0 1 2 0l1 1l8.64-8.64h0a1.38 1.38 0 0 1 .53-.28a.45.45 0 0 1 .11 0a1.56 1.56 0 0 1 .44-.09ZM20 31.26l-4.1 4.07l.23.23a2.82 2.82 0 0 0 4.07 0l2-2l-2.2-2.3Zm6.22-16.33l4.17 4.16M15.9 35.33l-4.08 4.06m10.42-5.87l4.09 4.04M13.99 25.27l5.99 5.99"/>`),
		g.Group(children),
	)
}