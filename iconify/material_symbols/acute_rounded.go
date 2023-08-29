package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcuteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 20q-3.35 0-5.675-2.325T7 12q0-3.325 2.325-5.663T15 4q3.325 0 5.663 2.337T23 12q0 3.35-2.337 5.675T15 20Zm1-8.4V9q0-.425-.288-.713T15 8q-.425 0-.713.288T14 9v3.025q0 .2.088.388t.212.312L16.575 15q.3.3.713.3T18 15q.3-.3.3-.712t-.3-.713L16 11.6ZM3 9q-.425 0-.713-.288T2 8q0-.425.288-.713T3 7h2q.425 0 .713.288T6 8q0 .425-.288.713T5 9H3Zm-1 4q-.425 0-.713-.288T1 12q0-.425.288-.713T2 11h3q.425 0 .713.288T6 12q0 .425-.288.713T5 13H2Zm1 4q-.425 0-.713-.288T2 16q0-.425.288-.713T3 15h2q.425 0 .713.288T6 16q0 .425-.288.713T5 17H3Z"/>`),
		g.Group(children),
	)
}