package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.486 4.768a7.25 7.25 0 1 0 7.399 9.51a6.25 6.25 0 0 1-7.398-9.51ZM3.25 12a8.75 8.75 0 0 1 10.074-8.65a.75.75 0 0 1 .336 1.342a4.75 4.75 0 1 0 5.83 7.499a.75.75 0 0 1 1.22.654A8.751 8.751 0 0 1 3.25 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}