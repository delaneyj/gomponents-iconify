package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProfileCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" d="M21 12a8.958 8.958 0 0 1-1.526 5.016A8.991 8.991 0 0 1 12 21a8.991 8.991 0 0 1-7.474-3.984A9 9 0 1 1 21 12Z"/><path fill="currentColor" d="M13.5 9a1.5 1.5 0 0 1-1.5 1.5v1A2.5 2.5 0 0 0 14.5 9h-1ZM12 10.5A1.5 1.5 0 0 1 10.5 9h-1a2.5 2.5 0 0 0 2.5 2.5v-1ZM10.5 9A1.5 1.5 0 0 1 12 7.5v-1A2.5 2.5 0 0 0 9.5 9h1ZM12 7.5A1.5 1.5 0 0 1 13.5 9h1A2.5 2.5 0 0 0 12 6.5v1ZM5.166 17.856l-.48-.142l-.077.261l.177.207l.38-.326Zm13.668 0l.38.326l.177-.207l-.078-.261l-.479.142ZM9 15.5h6v-1H9v1Zm0-1a4.502 4.502 0 0 0-4.313 3.214l.958.285A3.502 3.502 0 0 1 9 15.5v-1Zm3 6a8.48 8.48 0 0 1-6.455-2.97l-.759.652A9.48 9.48 0 0 0 12 21.5v-1Zm3-5a3.502 3.502 0 0 1 3.355 2.5l.958-.286A4.502 4.502 0 0 0 15 14.5v1Zm3.455 2.03A8.48 8.48 0 0 1 12 20.5v1a9.48 9.48 0 0 0 7.214-3.318l-.76-.651Z"/></g>`),
		g.Group(children),
	)
}