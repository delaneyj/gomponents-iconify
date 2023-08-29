package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsSit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23 22l7 5l-7 5V22z"/><path fill="currentColor" d="M26 3H6C4.346 3 3 4.346 3 6v20c0 1.654 1.346 3 3 3h11v-9h12V6c0-1.654-1.346-3-3-3ZM6 5h20a1 1 0 0 1 1 1v3H5V6a1 1 0 0 1 1-1Zm9 6v7H5v-7h10Zm0 16H6a1 1 0 0 1-1-1v-6h10v7Zm2-9v-7h10v7H17Z"/>`),
		g.Group(children),
	)
}