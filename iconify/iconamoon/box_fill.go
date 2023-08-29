package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a1 1 0 0 0-.707.293l-4 4A1 1 0 0 0 3 7.954m0 .054V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.008l-.001-.054a1 1 0 0 0-.292-.661l-4-4A1 1 0 0 0 16 3H8m9.586 4l-2-2H8.414l-2 2h11.172ZM7 12a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}