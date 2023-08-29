package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Size(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 6V4h8v16h-8v-2H8v-2H4V8h4V6h4Zm2 0h4v12h-4V6Zm-2 2h-2v8h2V8Zm-4 2v4H6v-4h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}