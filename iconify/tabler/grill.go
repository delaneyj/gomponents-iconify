package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 8H5a6 6 0 0 0 6 6h2a6 6 0 0 0 6-5.775V8zm-2 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm-2-6l1 2m-7-2l-3 6m9-2H7m8-13V4m-3 1V4M9 5V4"/>`),
		g.Group(children),
	)
}