package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopChromebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 20v-2h2V3h20v15h2v2H0Zm10-2h4v-1h-4v1Z"/>`),
		g.Group(children),
	)
}