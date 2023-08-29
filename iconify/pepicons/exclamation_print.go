package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><g opacity=".8"><path fill-rule="evenodd" d="M11 3a2 2 0 0 1 2 2v7a2 2 0 1 1-4 0V5a2 2 0 0 1 2-2Z" clip-rule="evenodd"/><path d="M13 17a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g><path fill-rule="evenodd" d="M10.25 2.25A.75.75 0 0 1 11 3v9a.75.75 0 0 1-1.5 0V3a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path d="M11.5 16.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/></g>`),
		g.Group(children),
	)
}