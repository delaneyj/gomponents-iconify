package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tikcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 5H5v4h4V5ZM3 3v8h8V3H3Zm16 2h-4v4h4V5Zm-6-2v8h8V3h-8ZM9 15H5v4h4v-4Zm-6-2v8h8v-8H3Z" clip-rule="evenodd"/><path d="M13 13h2v8h-2v-8Zm3 0h2v8h-2v-8Zm3 0h2v8h-2v-8Z"/></g>`),
		g.Group(children),
	)
}