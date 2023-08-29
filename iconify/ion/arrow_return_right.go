package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowReturnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M448 192l-128 96v-64H128v128h248c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H72c-4.4 0-8-3.6-8-8V168c0-4.4 3.6-8 8-8h248V96l128 96z" fill="currentColor"/>`),
		g.Group(children),
	)
}