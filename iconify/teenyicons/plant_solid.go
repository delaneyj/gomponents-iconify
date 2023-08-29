package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlantSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 4.5A4.5 4.5 0 0 1 11.5 0H15v3.5A4.5 4.5 0 0 1 10.5 8H8v7H7v-4H4.5A4.5 4.5 0 0 1 0 6.5V3h3.5c1.414 0 2.675.652 3.5 1.671V4.5Zm1.146 1.646l3-3l.708.708l-3 3l-.708-.708Zm-2 3.708l-3-3l.708-.708l3 3l-.708.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}