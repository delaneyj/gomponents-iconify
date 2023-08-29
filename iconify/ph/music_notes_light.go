package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M211.69 27.27a6 6 0 0 0-5.15-1.09l-128 32A6 6 0 0 0 74 64v114.11A34 34 0 1 0 86 204v-87.32l116-29v58.43A34 34 0 1 0 214 172V32a6 6 0 0 0-2.31-4.73ZM52 226a22 22 0 1 1 22-22a22 22 0 0 1-22 22Zm34-121.68V68.68l116-29v35.64ZM180 194a22 22 0 1 1 22-22a22 22 0 0 1-22 22Z"/>`),
		g.Group(children),
	)
}