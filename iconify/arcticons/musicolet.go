package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musicolet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.52 4.5a1.14 1.14 0 0 0-1.15 1.15v36.7a1.13 1.13 0 0 0 1.68 1l15.51-9.18L40.07 25a1.17 1.17 0 0 0 0-2l-15.51-9.17L9.05 4.65a1.2 1.2 0 0 0-.53-.15Zm12 14.78v11.26h0v0a3.26 3.26 0 0 1-3.49 3a3.26 3.26 0 0 1-3.49-3a3.26 3.26 0 0 1 3.49-3a4.19 4.19 0 0 1 3.5 1.63V15.72a5.16 5.16 0 0 0 .28 1.54a5.26 5.26 0 0 0 1 1.09a5.18 5.18 0 0 1 2 4.2"/>`),
		g.Group(children),
	)
}