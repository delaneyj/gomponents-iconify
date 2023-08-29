package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OperationIf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M12 13h2v10h-2z" fill="currentColor"/><path d="M12 9h2v2h-2z" fill="currentColor"/><path d="M23 11V9h-3a2 2 0 0 0-2 2v2h-2v2h2v8h2v-8h3v-2h-3v-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}