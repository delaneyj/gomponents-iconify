package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.25 6.47v19.06h25.5V6.47H3.25zm7.095 5.043l-4.33 1.926v-1l3.123-1.288v-.018L6.014 9.848v-1l4.33 1.927v.738zM16.04 14.6h-5.05v-.88h5.05v.88z"/>`),
		g.Group(children),
	)
}