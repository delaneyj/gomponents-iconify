package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlbumOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4h10a2 2 0 0 1 2 2v10m-.581 3.41c-.362.364-.864.59-1.419.59H6a2 2 0 0 1-2-2V6c0-.552.224-1.052.585-1.413"/><path d="M12 4v4m1.503 1.497L14 9l2 2V4M3 3l18 18"/></g>`),
		g.Group(children),
	)
}