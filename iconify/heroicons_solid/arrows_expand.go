package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h4a1 1 0 0 1 0 2H6.414l2.293 2.293a1 1 0 0 1-1.414 1.414L5 6.414V8a1 1 0 0 1-2 0V4Zm9 1a1 1 0 1 1 0-2h4a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V6.414l-2.293 2.293a1 1 0 1 1-1.414-1.414L13.586 5H12Zm-9 7a1 1 0 1 1 2 0v1.586l2.293-2.293a1 1 0 0 1 1.414 1.414L6.414 15H8a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1v-4Zm13-1a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 1 1 0-2h1.586l-2.293-2.293a1 1 0 0 1 1.414-1.414L15 13.586V12a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}