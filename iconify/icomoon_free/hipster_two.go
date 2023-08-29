package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HipsterTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2zM5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm3.497 6.615A1.507 1.507 0 0 1 8 9.5a1.491 1.491 0 0 1-.497 1.115C6.151 11.842 3 10.586 3 9c.969.625 1.726.153 2.439-.561a1.5 1.5 0 0 1 2.56 1.06a1.5 1.5 0 0 1 2.56-1.06c.713.714 1.471 1.186 2.439.561c0 1.586-3.151 2.842-4.503 1.615z"/>`),
		g.Group(children),
	)
}