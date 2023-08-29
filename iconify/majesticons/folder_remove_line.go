package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderRemoveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-5.93a3 3 0 0 1-2.496-1.336l-.812-1.219A1 1 0 0 0 8.93 5H5zM2 6a3 3 0 0 1 3-3h3.93a3 3 0 0 1 2.496 1.336l.812 1.219A1 1 0 0 0 13.07 6H19a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6z"/><path d="M8 14a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1z"/></g>`),
		g.Group(children),
	)
}