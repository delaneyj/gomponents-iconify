package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulbOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M304 384v-24c0-29 31.54-56.43 52-76c28.84-27.57 44-64.61 44-108c0-80-63.73-144-144-144a143.6 143.6 0 0 0-144 144c0 41.84 15.81 81.39 44 108c20.35 19.21 52 46.7 52 76v24m16 96h64m-80-48h96m-48-48V256"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M294 240s-21.51 16-38 16s-38-16-38-16"/>`),
		g.Group(children),
	)
}