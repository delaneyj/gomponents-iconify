package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassMartini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20h-5v-5.06A9 9 0 0 0 21 6a8.72 8.72 0 0 0-.67-3.39a1 1 0 0 0-.22-.32L20 2.21a.92.92 0 0 0-.21-.13a.94.94 0 0 0-.28-.08H4.5a.94.94 0 0 0-.29.06A2.12 2.12 0 0 0 4 2.2l-.12.09a1 1 0 0 0-.22.32A8.72 8.72 0 0 0 3 6a9 9 0 0 0 8 8.94V20H6a1 1 0 0 0 0 2h12a1 1 0 0 0 0-2ZM5 6a6.91 6.91 0 0 1 .29-2h13.42A6.91 6.91 0 0 1 19 6A7 7 0 0 1 5 6Z"/>`),
		g.Group(children),
	)
}