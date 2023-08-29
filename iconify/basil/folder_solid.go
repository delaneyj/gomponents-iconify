package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.585 17.155c.443-2.42.47-4.9.082-7.33l-.064-.4a2.158 2.158 0 0 0-2.13-1.819H9.758a.06.06 0 0 1-.059-.06c0-.992-.804-1.796-1.797-1.796h-2.29a2.18 2.18 0 0 0-2.164 1.92l-.273 2.269a23.73 23.73 0 0 0 .217 7.094a2.128 2.128 0 0 0 1.942 1.74l1.514.11c3.43.245 6.874.245 10.304 0l1.638-.118a1.968 1.968 0 0 0 1.795-1.61Z"/>`),
		g.Group(children),
	)
}