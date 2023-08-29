package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.6 4H1.4C.629 4 0 4.629 0 5.4v9.2c0 .769.629 1.4 1.399 1.4h17.2c.77 0 1.4-.631 1.4-1.4V5.4C20 4.629 19.369 4 18.6 4zM11 6h2v2h-2V6zm3 3v2h-2V9h2zM8 6h2v2H8V6zm3 3v2H9V9h2zM5 6h2v2H5V6zm3 3v2H6V9h2zM2 6h2v2H2V6zm3 3v2H3V9h2zm-1 5H2v-2h2v2zm11 0H5v-2h10v2zm3 0h-2v-2h2v2zm-3-3V9h2v2h-2zm3-3h-4V6h4v2z"/>`),
		g.Group(children),
	)
}