package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuyaSmart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.43 12.783v19.862a5.855 5.855 0 0 0 5.855 5.855h2.706M14.545 18.888h11.521m4.706 0a4.705 4.705 0 0 0-4.706-4.705m9.389 4.705A9.388 9.388 0 0 0 26.066 9.5"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}