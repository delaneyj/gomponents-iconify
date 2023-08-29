package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kanban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3H3v18h18V3zM5 19V5h14v14H5zM9 7H7v8h2V7zm2 0h2v4h-2V7zm6 0h-2v10h2V7z"/>`),
		g.Group(children),
	)
}