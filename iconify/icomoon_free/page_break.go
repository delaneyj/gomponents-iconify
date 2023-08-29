package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 8h2v1H0zm3 0h3v1H3zm4 0h2v1H7zm3 0h3v1h-3zm4 0h2v1h-2zm-.25-8L14 7H2l.25-7h.5L3 6h10l.25-6zM2.25 16L2 10h12l-.25 6h-.5L13 11H3l-.25 5z"/>`),
		g.Group(children),
	)
}