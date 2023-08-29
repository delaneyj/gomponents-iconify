package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowEnterTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M21 4a1 1 0 0 1 .993.883L22 5v6.5a3.5 3.5 0 0 1-3.308 3.495L18.5 15H5.415l3.292 3.293a1 1 0 0 1 .083 1.32l-.083.094a1 1 0 0 1-1.32.083l-.094-.083l-5-5a1.008 1.008 0 0 1-.097-.112l-.071-.11l-.054-.114l-.035-.105l-.025-.118l-.007-.058L2 14l.003-.075l.017-.126l.03-.111l.044-.111l.052-.098l.064-.092l.083-.094l5-5a1 1 0 0 1 1.497 1.32l-.083.094L5.415 13H18.5a1.5 1.5 0 0 0 1.493-1.356L20 11.5V5a1 1 0 0 1 1-1z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}