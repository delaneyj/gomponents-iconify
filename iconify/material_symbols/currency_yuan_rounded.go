package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyYuanRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-.425 0-.713-.288T11 20v-6H7q-.425 0-.713-.288T6 13q0-.425.288-.713T7 12h3.725l-4.75-7.45q-.325-.5-.037-1.025T6.825 3q.275 0 .5.125t.35.35L12 10.3l4.325-6.825q.125-.225.35-.35t.5-.125q.6 0 .888.525t-.038 1.025L13.275 12H17q.425 0 .713.288T18 13q0 .425-.288.713T17 14h-4v6q0 .425-.288.713T12 21Z"/>`),
		g.Group(children),
	)
}