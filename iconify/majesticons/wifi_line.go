package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2.041 10.643A12.97 12.97 0 0 1 12 6a12.97 12.97 0 0 1 9.959 4.643a1 1 0 1 0 1.53-1.286A14.97 14.97 0 0 0 12 4A14.97 14.97 0 0 0 .51 9.357a1 1 0 0 0 1.531 1.286zM5.573 13.7A8.97 8.97 0 0 1 12 11a8.97 8.97 0 0 1 6.427 2.7a1 1 0 0 0 1.428-1.4A10.97 10.97 0 0 0 12 9a10.97 10.97 0 0 0-7.856 3.3a1 1 0 1 0 1.429 1.4zm3.663 3.133A4.972 4.972 0 0 1 12 16c1.024 0 1.973.307 2.764.833a1 1 0 1 0 1.107-1.666A6.972 6.972 0 0 0 12 14c-1.43 0-2.762.43-3.871 1.167a1 1 0 1 0 1.107 1.666zM12 18a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H12z"/></g>`),
		g.Group(children),
	)
}