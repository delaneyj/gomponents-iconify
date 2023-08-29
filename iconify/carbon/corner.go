package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Corner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 9H14V6H6v8h3v14h2V14h3v-3h14Zm-16 3H8V8h4Z"/>`),
		g.Group(children),
	)
}