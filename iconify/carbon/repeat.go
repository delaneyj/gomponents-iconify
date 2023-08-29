package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 6h20.172l-3.586-3.586L24 1l6 6l-6 6l-1.414-1.414L26.172 8H6v7H4V8a2.002 2.002 0 0 1 2-2zm3.414 14.414L5.828 24H26v-7h2v7a2.002 2.002 0 0 1-2 2H5.828l3.586 3.586L8 31l-6-6l6-6z"/>`),
		g.Group(children),
	)
}