package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderMoveTo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18 13l-1.41 1.41L19.17 17H10v2h9.17l-2.58 2.59L18 23l5-5l-5-5z"/><path fill="currentColor" d="m11.172 6l3.414 3.414l.586.586H28v16H4V6h7.172m0-2H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2H16l-3.414-3.414A2 2 0 0 0 11.172 4Z"/>`),
		g.Group(children),
	)
}