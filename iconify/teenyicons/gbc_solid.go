package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GbcSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 6V3h5v2.5a.5.5 0 0 1-.5.5H5Z"/><path fill="currentColor" fill-rule="evenodd" d="M2 1.5A1.5 1.5 0 0 1 3.5 0h8A1.5 1.5 0 0 1 13 1.5v10A3.5 3.5 0 0 1 9.5 15h-6A1.5 1.5 0 0 1 2 13.5v-12Zm2.5.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h5A1.5 1.5 0 0 0 11 5.5v-3a.5.5 0 0 0-.5-.5h-6ZM5 8v1H4v1h1v1h1v-1h1V9H6V8H5Zm5 0v1h1V8h-1Zm-1 3v-1h1v1H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}