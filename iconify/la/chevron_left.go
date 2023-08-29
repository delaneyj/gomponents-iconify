package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.75 2.594l-.719.687l-12 12l-.687.719l.687.719l12 12l.719.687l.719-.687l3.593-3.625l.688-.688l-.688-.718L16.375 16l7.688-7.688l.687-.718l-.688-.688l-3.593-3.625zm0 2.844l2.188 2.187l-7.688 7.656l-.719.719l.719.719l7.688 7.656l-2.188 2.188L9.187 16z"/>`),
		g.Group(children),
	)
}