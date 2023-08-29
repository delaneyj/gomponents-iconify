package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewModule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.675 11.5v-6H21v6h-5.325Zm-6.325 0v-6h5.325v6H9.35Zm-6.325 0v-6H8.35v6H3.025Zm0 7v-6H8.35v6H3.025Zm6.325 0v-6h5.325v6H9.35Zm6.325 0v-6H21v6h-5.325Z"/>`),
		g.Group(children),
	)
}