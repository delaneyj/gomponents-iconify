package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OffClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22A9.99 9.99 0 0 1 2 12v-.2a10.005 10.005 0 0 1 17.074-6.874A10 10 0 0 1 12 22Zm0-8.59L14.59 16L16 14.59L13.41 12L16 9.41L14.59 8L12 10.59L9.41 8L8 9.41L10.59 12L8 14.59L9.41 16L12 13.411v-.001Z"/>`),
		g.Group(children),
	)
}