package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.75 2.594l-.72.687l-12 12l-.686.72l.687.72l12 12l.72.686l.72-.687l3.593-3.626l.687-.688l-.688-.718L16.375 16l7.688-7.688l.687-.718l-.688-.688L20.47 3.28l-.72-.686zm0 2.844l2.188 2.187l-7.688 7.656l-.72.72l.72.72l7.688 7.655l-2.188 2.188L9.187 16L19.75 5.437z"/>`),
		g.Group(children),
	)
}