package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 13h-2v6h2a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M28 6H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2ZM15 21h-2v-4h-3v4H8V11h2v4h3v-4h2Zm10-3a3 3 0 0 1-3 3h-4V11h4a3 3 0 0 1 3 3Z"/>`),
		g.Group(children),
	)
}