package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkerStrokedEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.486 11l-.365-.446c-.7-.858-3.544-4.739-3.544-6.638A3.726 3.726 0 0 1 5.132.026q.167-.008.333 0a3.726 3.726 0 0 1 3.888 3.557q.007.166 0 .331c0 1.6-1.806 4.268-3.38 6.415zM5.465.916a2.817 2.817 0 0 0-3 3c0 1.268 1.883 4.161 2.987 5.62c.935-1.282 3.011-4.217 3.011-5.62a2.817 2.817 0 0 0-3-3z" fill="currentColor"/>`),
		g.Group(children),
	)
}