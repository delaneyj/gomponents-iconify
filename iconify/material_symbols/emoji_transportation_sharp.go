package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiTransportationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22v-6.35L11.625 11h8.75L22 15.65V22h-2v-1.5h-8V22h-2Zm2-7.5h8l-.7-2h-6.6l-.7 2Zm1 4q.425 0 .713-.288T14 17.5q0-.425-.288-.713T13 16.5q-.425 0-.713.288T12 17.5q0 .425.288.713T13 18.5Zm6 0q.425 0 .713-.288T20 17.5q0-.425-.288-.713T19 16.5q-.425 0-.713.288T18 17.5q0 .425.288.713T19 18.5ZM6 14v-2h2v2H6Zm5-6V6h2v2h-2ZM6 18v-2h2v2H6Zm0 4v-2h2v2H6Zm-4 0V8h5V2h10v7h-2V4H9v6H4v12H2Z"/>`),
		g.Group(children),
	)
}