package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyYenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-.425 0-.713-.288T11 20v-3H7q-.425 0-.713-.288T6 16q0-.425.288-.713T7 15h4v-2H7q-.425 0-.713-.288T6 12q0-.425.288-.713T7 11h3.075l-4.1-6.475q-.325-.5-.037-1.012T6.825 3q.275 0 .5.125t.35.35L12 10.3l4.325-6.825q.125-.225.35-.35t.5-.125q.6 0 .888.525t-.038 1.025l-4.1 6.45H17q.425 0 .713.288T18 12q0 .425-.288.713T17 13h-4v2h4q.425 0 .713.288T18 16q0 .425-.288.713T17 17h-4v3q0 .425-.288.713T12 21Z"/>`),
		g.Group(children),
	)
}