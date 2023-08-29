package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7 3.093l-5 5V8.8l5 5l.707-.707l-4.146-4.147H14v-1H3.56L7.708 3.8L7 3.093z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}