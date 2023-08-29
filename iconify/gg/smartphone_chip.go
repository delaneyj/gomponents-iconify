package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneChip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9 22a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M9 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM9 22a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm7-7a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm0-4a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm-1-5a1 1 0 1 1 0 2a1 1 0 0 1 0-2ZM2 15a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm1-3a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm14-2H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1ZM7 4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H7Zm7 6h-4v4h4v-4ZM8 8v8h8V8H8Z"/></g>`),
		g.Group(children),
	)
}