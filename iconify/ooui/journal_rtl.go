package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JournalRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 0v20h1.5c.8 0 1.5-.7 1.5-1.5v-17c0-.8-.7-1.5-1.5-1.5zM2 18c0 1.1.9 2 2 2h10V0H4C2.9 0 2 .9 2 2zM4 5h8v1H4zm3 2h5v1H7z"/>`),
		g.Group(children),
	)
}