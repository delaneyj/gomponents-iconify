package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncProblem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-2h2.75l-.4-.35q-1.225-1.225-1.788-2.663T3 12.05q0-2.775 1.663-4.938T9 4.25v2.1Q7.2 7 6.1 8.562T5 12.05q0 1.125.425 2.188T6.75 16.2l.25.25V14h2v6H3Zm9-3q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15q.425 0 .713.288T13 16q0 .425-.288.713T12 17Zm-1-4V7h2v6h-2Zm4 6.75v-2.1q1.8-.65 2.9-2.212T19 11.95q0-1.125-.425-2.187T17.25 7.8L17 7.55V10h-2V4h6v2h-2.75l.4.35q1.225 1.225 1.788 2.663T21 11.95q0 2.775-1.663 4.938T15 19.75Z"/>`),
		g.Group(children),
	)
}