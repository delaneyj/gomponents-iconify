package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilitaryCamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 28V17.52a2.003 2.003 0 0 0-.853-1.639L17 9.48V8h6V2h-8v7.48l-9.147 6.402A2.002 2.002 0 0 0 5 17.52V28H2v2h28v-2ZM17 4h4v2h-4ZM7 17.52l9-6.3l9 6.3V28h-8v-8h-2v8H7Z"/>`),
		g.Group(children),
	)
}