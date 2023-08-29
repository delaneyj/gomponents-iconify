package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Style(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 21v-8h8v8h-8Zm2-6h4v4h-4v-4ZM3 11V3h8v8H3Zm2-6h4v4H5V5Z" clip-rule="evenodd"/><path d="M18 6v6h-2V8h-4V6h6Zm-6 12H6v-6h2v4h4v2Z"/></g>`),
		g.Group(children),
	)
}