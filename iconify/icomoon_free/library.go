package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Library(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 15v-1h-1V8h1V7h-3v1h1v6h-3V8h1V7H9v1h1v6H7V8h1V7H5v1h1v6H3V8h1V7H1v1h1v6H1v1H0v1h17v-1h-1zM8 0h1l8 5v1H0V5l8-5z"/>`),
		g.Group(children),
	)
}