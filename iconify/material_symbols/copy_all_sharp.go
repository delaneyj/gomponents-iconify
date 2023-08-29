package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyAllSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18V2h13v16H7Zm2-2h9V4H9v12Zm-6-1h2v-2H3v2Zm0-3.5h2v-2H3v2ZM10 22h2v-2h-2v2Zm-7-3.5h2v-2H3v2ZM3 22h2v-2H3v2Zm3.5 0h2v-2h-2v2Zm7 0h2v-2h-2v2ZM3 8h2V6H3v2Z"/>`),
		g.Group(children),
	)
}