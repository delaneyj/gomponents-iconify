package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 12a3 3 0 0 1 3-3h16a3 3 0 0 1 2.273 4.958c-1.708 1.982-4.27 5.601-6.558 9.802c-2.303 4.229-4.18 8.783-4.746 12.672a3 3 0 0 1-5.938-.864c.715-4.91 2.97-10.19 5.414-14.678a75.281 75.281 0 0 1 3.57-5.89H16a3 3 0 0 1-3-3Zm3-1a1 1 0 1 0 0 2h11.922a1 1 0 0 1 .815 1.58a72.375 72.375 0 0 0-4.535 7.267c-2.397 4.401-4.527 9.439-5.192 14.01a1 1 0 1 0 1.98.287c.615-4.23 2.618-9.026 4.968-13.34c2.345-4.307 4.986-8.046 6.8-10.151A1 1 0 0 0 32 11H16Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}