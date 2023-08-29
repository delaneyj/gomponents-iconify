package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookCheckSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 3.5A4.5 4.5 0 0 0 4 8v10h.035A3.5 3.5 0 0 0 7.5 21H19a1 1 0 0 0 1-1V5a1.5 1.5 0 0 0-1.5-1.5h-10Zm-1 12h11v4h-11a2 2 0 1 1 0-4Zm8.02-6.959a.75.75 0 1 0-1.04-1.082L11.876 9.96l-1.357-1.3a.75.75 0 1 0-1.038 1.082l1.875 1.8a.75.75 0 0 0 1.038 0l3.125-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}