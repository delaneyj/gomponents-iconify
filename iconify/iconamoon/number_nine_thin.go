package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.576 19.734a.5.5 0 1 0 .848.532l-.848-.532Zm6.074-7.796a.5.5 0 1 0-.847-.532l.847.532ZM7.5 9A4.5 4.5 0 0 1 12 4.5v-1A5.5 5.5 0 0 0 6.5 9h1ZM12 4.5A4.5 4.5 0 0 1 16.5 9h1A5.5 5.5 0 0 0 12 3.5v1ZM16.5 9a4.5 4.5 0 0 1-4.5 4.5v1A5.5 5.5 0 0 0 17.5 9h-1ZM12 13.5A4.5 4.5 0 0 1 7.5 9h-1a5.5 5.5 0 0 0 5.5 5.5v-1Zm-.576 6.766l5.226-8.328l-.847-.532l-5.227 8.328l.848.532Z"/>`),
		g.Group(children),
	)
}