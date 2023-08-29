package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoEntryPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 11.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Z" opacity=".8"/><path d="M10 3.5a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13ZM2.5 10a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Z"/><path d="M15.304 4.697a.5.5 0 0 1 0 .707l-9.9 9.9a.5.5 0 1 1-.707-.707l9.9-9.9a.5.5 0 0 1 .707 0Z"/></g>`),
		g.Group(children),
	)
}