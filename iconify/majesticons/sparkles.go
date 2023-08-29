package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparkles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 3a1 1 0 0 0-2 0v1H3a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0V6h1a1 1 0 0 0 0-2H6V3zm7-1a1 1 0 0 1 1 1c0 3.344 1.148 5.296 2.514 6.43C17.918 10.598 19.672 11 21 11a1 1 0 1 1 0 2c-3.26 0-4.924 1.324-5.838 2.881C14.2 17.524 14 19.556 14 21a1 1 0 1 1-2 0c0-3.344-1.148-5.296-2.514-6.43C8.082 13.402 6.328 13 5 13a1 1 0 1 1 0-2c3.26 0 4.924-1.324 5.838-2.881C11.8 6.476 12 4.444 12 3a1 1 0 0 1 1-1z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}