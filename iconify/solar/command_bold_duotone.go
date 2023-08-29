package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommandBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16 16h3a3 3 0 1 1-3 3.001V16ZM5 16l3 .001v3a3 3 0 1 1-3-3Z"/><path fill-rule="evenodd" d="M19 8h-3V5a3 3 0 1 1 3 3ZM8 8V5a3 3 0 1 0-3 3h3Z" clip-rule="evenodd"/><path d="M16 8H8v8h8V8Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}