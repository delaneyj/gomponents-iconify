package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRupeeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.15 21q-.2 0-.388-.075t-.337-.225L7.4 14.425q-.175-.175-.288-.475T7 13.4q0-.575.413-.988T8.4 12h2.1q1.325 0 2.288-.863T13.95 9H7q-.425 0-.713-.288T6 8q0-.425.288-.713T7 7h6.65q-.425-.875-1.263-1.438T10.5 5H7q-.425 0-.713-.288T6 4q0-.425.288-.713T7 3h10q.425 0 .713.288T18 4q0 .425-.288.713T17 5h-2.25q.35.425.625.925T15.8 7H17q.425 0 .713.288T18 8q0 .425-.288.713T17 9h-1.025q-.2 2.125-1.75 3.563T10.5 14h-.725l5.1 5.3q.45.475.188 1.088T14.15 21Z"/>`),
		g.Group(children),
	)
}