package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3.936L4.361 8.168L6.45 23.84L16 29.145l9.55-5.307l2.089-15.67L16 3.936zm0 2.128l9.443 3.434l-1.744 13.08L16 26.855l-7.7-4.277l-1.743-13.08L16 6.064zM16 8l-6 14h2l1.29-3h5.42L20 22h2L16 8zm0 4.55l.02.06l.7 1.75L17.85 17h-3.7l1.13-2.64l.7-1.75l.02-.06z"/>`),
		g.Group(children),
	)
}