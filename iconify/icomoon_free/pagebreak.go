package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pagebreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 6V0h12v6h-1V1H5v5zm12 3v7H4V9h1v6h10V9zM8 7h2v1H8zM5 7h2v1H5zm6 0h2v1h-2zm3 0h2v1h-2zM0 4.5l3 3l-3 3z"/>`),
		g.Group(children),
	)
}