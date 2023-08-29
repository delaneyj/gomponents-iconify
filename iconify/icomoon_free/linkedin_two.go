package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 6h2.767v1.418h.04C9.192 6.727 10.134 6 11.539 6C14.46 6 15 7.818 15 10.183V15h-2.885v-4.27c0-1.018-.021-2.329-1.5-2.329c-1.502 0-1.732 1.109-1.732 2.255V15H6V6zM1 6h3v9H1V6zm3-2.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 4 3.5z"/>`),
		g.Group(children),
	)
}