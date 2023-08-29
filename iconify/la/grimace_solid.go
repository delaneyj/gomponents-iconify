package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrimaceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm-4.5 7a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3zm9 0a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3zM12 17c-1.654 0-3 1.346-3 3s1.346 3 3 3h8c1.654 0 3-1.346 3-3s-1.346-3-3-3h-8zm0 2h1v2h-1a1 1 0 0 1 0-2zm3 0h2v2h-2v-2zm4 0h1a1 1 0 0 1 0 2h-1v-2z"/>`),
		g.Group(children),
	)
}