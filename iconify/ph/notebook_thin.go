package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M180 112a4 4 0 0 1-4 4h-64a4 4 0 0 1 0-8h64a4 4 0 0 1 4 4Zm-4 28h-64a4 4 0 0 0 0 8h64a4 4 0 0 0 0-8Zm44-92v160a12 12 0 0 1-12 12H48a12 12 0 0 1-12-12V48a12 12 0 0 1 12-12h160a12 12 0 0 1 12 12ZM48 212h28V44H48a4 4 0 0 0-4 4v160a4 4 0 0 0 4 4ZM212 48a4 4 0 0 0-4-4H84v168h124a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}