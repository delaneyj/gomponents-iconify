package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.5 1.75a.75.75 0 0 0-1.5 0v20a.75.75 0 0 0 1.5 0v-20Z" clip-rule="evenodd" opacity=".5"/><path d="m13.349 3.79l-.204-.082a8.677 8.677 0 0 0-4.924-.452L6.5 3.6v10l1.72-.344a8.677 8.677 0 0 1 4.925.452a8.677 8.677 0 0 0 5.327.361l.214-.053a1.404 1.404 0 0 0 1.064-1.362V5.287a1.2 1.2 0 0 0-1.49-1.164a7.999 7.999 0 0 1-4.911-.334Z"/></g>`),
		g.Group(children),
	)
}