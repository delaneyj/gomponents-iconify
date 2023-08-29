package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiTransportation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21.35v-5.7l1.4-4q.125-.275.363-.463T12.35 11h7.3q.35 0 .588.188t.362.462l1.4 4v5.7q0 .275-.188.463T21.35 22h-.7q-.275 0-.462-.188T20 21.35v-.85h-8v.85q0 .275-.188.463T11.35 22h-.7q-.275 0-.462-.188T10 21.35Zm2-6.85h8l-.7-2h-6.6l-.7 2Zm1 4q.425 0 .713-.288T14 17.5q0-.425-.288-.713T13 16.5q-.425 0-.713.288T12 17.5q0 .425.288.713T13 18.5Zm6 0q.425 0 .713-.288T20 17.5q0-.425-.288-.713T19 16.5q-.425 0-.713.288T18 17.5q0 .425.288.713T19 18.5ZM6 14v-2h2v2H6Zm5-6V6h2v2h-2ZM6 18v-2h2v2H6Zm0 4v-2h2v2H6Zm-4 0V8h5V2h10v7h-2V4H9v6H4v12H2Z"/>`),
		g.Group(children),
	)
}