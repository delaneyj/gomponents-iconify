package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20v-2h2.75l-.4-.35q-1.225-1.225-1.788-2.663T4 12.05q0-2.775 1.663-4.938T10 4.25v2.1Q8.2 7 7.1 8.562T6 12.05q0 1.125.425 2.188T7.75 16.2l.25.25V14h2v6H4Zm10-.25v-2.1q1.8-.65 2.9-2.212T18 11.95q0-1.125-.425-2.187T16.25 7.8L16 7.55V10h-2V4h6v2h-2.75l.4.35q1.225 1.225 1.788 2.663T20 11.95q0 2.775-1.663 4.938T14 19.75Z"/>`),
		g.Group(children),
	)
}