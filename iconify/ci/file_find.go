package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileFind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7c.265 0 .52.105.707.293l6 6A.991.991 0 0 1 20 9v11a2 2 0 0 1-2 2ZM6 4v16h10.586l-2.566-2.566A3.941 3.941 0 0 1 12 18a4.044 4.044 0 1 1 3.434-1.98L18 18.588V9.414L12.586 4H6Zm6 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}