package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkspaceImport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 6v5H17V6h10m0-2H17a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zM12 25H6v-8h6v-2H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h6zm18-5H16.828l2.586-2.586L18 16l-5 5l5 5l1.414-1.414L16.828 22H30v-2zM11 6v5H6V6h5m0-2H6a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}