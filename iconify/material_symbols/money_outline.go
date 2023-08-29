package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 16h3q.425 0 .713-.288T19 15V9q0-.425-.288-.713T18 8h-3q-.425 0-.713.288T14 9v6q0 .425.288.713T15 16Zm1-2v-4h1v4h-1Zm-7 2h3q.425 0 .713-.288T13 15V9q0-.425-.288-.713T12 8H9q-.425 0-.713.288T8 9v6q0 .425.288.713T9 16Zm1-2v-4h1v4h-1Zm-5 2h2V8H5v8Zm-3 4V4h20v16H2ZM4 6v12V6Zm0 12h16V6H4v12Z"/>`),
		g.Group(children),
	)
}