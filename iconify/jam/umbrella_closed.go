package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 13.418L0 15L6.234 1.642a1 1 0 1 1 1.533 0L14 15l-6-1.582V17a1 1 0 0 0 2 0a1 1 0 0 1 2 0a3 3 0 0 1-6 0v-3.582zM7 4.73l-3.383 7.249L7 11.086l3.383.892L7 4.729z"/>`),
		g.Group(children),
	)
}