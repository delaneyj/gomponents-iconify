package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.978 14.895l.988-1.829c-4.372 0-10.502.51-12.888-1.036C1.028 11.524 1.042 11 1.042 11s.68 3.895 3.036 3.895h11.9zm-3.961-3.905L12.014 2l4.899 8.99h-4.896zm-1.002.032V-.052L6 10.253l5.015.77z"/>`),
		g.Group(children),
	)
}