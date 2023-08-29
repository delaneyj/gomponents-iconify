package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m5.5 15.5l10.3 5.947v-5.6l9.7 5.6V9.552l-9.7 5.6v-5.6z"/>`),
		g.Group(children),
	)
}