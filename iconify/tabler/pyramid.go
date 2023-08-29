package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.105 21.788a1.994 1.994 0 0 0 1.789 0l8.092-4.054c.538-.27.718-.951.385-1.452l-8.54-13.836a.999.999 0 0 0-1.664 0l-8.54 13.836a1.005 1.005 0 0 0 .386 1.452l8.092 4.054zM12 2v20"/>`),
		g.Group(children),
	)
}