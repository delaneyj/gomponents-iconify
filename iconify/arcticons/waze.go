package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Waze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.47 39.38a17.4 17.4 0 1 0-22-16.78a17.16 17.16 0 0 0 1.34 6.7h0c.08.21.19.41.29.62c1 2.41-.29 3.47-.29 3.47L2.77 40H11m7.14 0h5.55M28 18.44a3 3 0 0 1 3-3h0a3 3 0 0 1 3 3h0a3 3 0 0 1-3 3h0a3 3 0 0 1-3-3Zm-10.26 0a3 3 0 0 1 3-3h0a3 3 0 0 1 3 3h0a3 3 0 0 1-3 3h0a3 3 0 0 1-3-3Zm-.25 5.76a8.79 8.79 0 0 0 16.61 0"/><circle cx="14.52" cy="41.43" r="3.85" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="27.21" cy="41.43" r="3.85" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}