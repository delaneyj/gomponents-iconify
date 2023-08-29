package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDeleteSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.668 9.826c.387 2.43.36 4.909-.083 7.33a1.968 1.968 0 0 1-1.795 1.609l-1.638.117c-3.43.245-6.874.245-10.304 0l-1.514-.108a2.128 2.128 0 0 1-1.942-1.74a23.73 23.73 0 0 1-.217-7.095l.273-2.27A2.18 2.18 0 0 1 5.612 5.75h2.291c.993 0 1.797.804 1.797 1.797c0 .033.027.06.06.06h8.712c1.06 0 1.964.77 2.131 1.817l.064.402ZM9.702 11.202a.75.75 0 0 1 1.06 0L12 12.439l1.237-1.237a.75.75 0 0 1 1.061 1.06L13.061 13.5l1.237 1.237a.75.75 0 0 1-1.06 1.061L12 14.561l-1.237 1.237a.75.75 0 1 1-1.061-1.06l1.237-1.238l-1.237-1.237a.75.75 0 0 1 0-1.061Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}