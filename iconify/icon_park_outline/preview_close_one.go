package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviewCloseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9.858 18C6.238 21 4 24 4 24s8.954 12 20 12c1.37 0 2.708-.185 4-.508M20.032 12.5c1.282-.318 2.61-.5 3.968-.5c11.046 0 20 12 20 12s-2.239 3-5.858 6m-17.828-9.379a5 5 0 0 0 7.186 6.95M42 42L6 6"/>`),
		g.Group(children),
	)
}