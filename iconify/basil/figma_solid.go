package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.555 3.305A2.75 2.75 0 0 1 8.5 2.5h2.75V8H8.5a2.75 2.75 0 0 1-1.945-4.695ZM15.5 8h-2.75V2.5h2.75a2.75 2.75 0 0 1 0 5.5Zm0 1.5a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5Zm-8.945 7.805A2.75 2.75 0 0 1 8.5 16.5h2.75v2.75a2.75 2.75 0 1 1-4.695-1.945ZM8.5 9.5a2.75 2.75 0 1 0 0 5.5h2.75V9.5H8.5Z"/>`),
		g.Group(children),
	)
}