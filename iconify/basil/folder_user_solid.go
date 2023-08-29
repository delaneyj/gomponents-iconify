package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderUserSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.188 16.084a.224.224 0 0 1-.084.334a4.516 4.516 0 0 0-2.265 2.366a.422.422 0 0 1-.377.267a72.174 72.174 0 0 1-6.614-.169l-1.514-.108a2.128 2.128 0 0 1-1.942-1.74a23.73 23.73 0 0 1-.217-7.095l.273-2.27A2.18 2.18 0 0 1 4.612 5.75h2.291c.993 0 1.797.804 1.797 1.797c0 .033.027.06.06.06h8.712c1.06 0 1.964.77 2.131 1.817l.064.402c.042.26.079.522.111.784c.02.165-.16.28-.31.211a3.5 3.5 0 0 0-4.28 5.262Z"/><path fill="currentColor" d="M18 12a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 8.5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}