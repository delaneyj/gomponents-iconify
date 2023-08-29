package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductHunt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 19a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm-9-7a9 9 0 1 1 18 0a9 9 0 0 1-18 0Zm6 4V8h4a3 3 0 1 1 0 6h-2v2H9Zm5-5a1 1 0 0 1-1 1h-2v-2h2a1 1 0 0 1 1 1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}