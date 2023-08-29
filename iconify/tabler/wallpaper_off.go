package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallpaperOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 6h8a2 2 0 0 1 2 2v8m-.58 3.409A2 2 0 0 1 18 20H6"/><path d="M4 18a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M8 18V8M4.573 4.598A2.003 2.003 0 0 0 4 6v12M3 3l18 18"/></g>`),
		g.Group(children),
	)
}