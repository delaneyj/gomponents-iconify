package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nrk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.387 18.1v11.8m10.587-11.8v11.8m6.831-11.8v11.8M14.54 18.1l4.281 11.8m17.587-11.8l-4.45 5.9l4.655 5.9"/><circle cx="25.389" cy="18.351" r=".75" fill="currentColor"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}