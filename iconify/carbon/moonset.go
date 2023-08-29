package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moonset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 28h28v2H2zm24-2h-2a7.987 7.987 0 0 0-2.037-5.333l1.49-1.335A9.983 9.983 0 0 1 26 26zm-10 0h-2a9.927 9.927 0 0 1 3.754-7.804A7.89 7.89 0 0 0 16 18a8.01 8.01 0 0 0-8 8H6a10.011 10.011 0 0 1 10-10a9.892 9.892 0 0 1 4.446 1.052a1 1 0 0 1 0 1.79A7.957 7.957 0 0 0 16 26zm0-12l-5-5l1.41-1.41L15 10.17V2h2v8.17l2.59-2.58L21 9l-5 5z"/>`),
		g.Group(children),
	)
}