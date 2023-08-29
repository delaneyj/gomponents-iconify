package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataSaverOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.15q-1.35-1.362-2.137-3.187T2 12q0-3.925 2.6-6.75t6.4-3.2v3q-2.575.35-4.288 2.313T5 12q0 2.9 2.05 4.95T12 19q1.65 0 3.088-.7t2.412-1.9l2.6 1.5q-1.35 1.875-3.475 2.988T12 22Zm-1-6v-3H8v-2h3V8h2v3h3v2h-3v3h-2Zm10.15.05l-2.6-1.5q.225-.6.337-1.238T19 12q0-2.675-1.713-4.638T13 5.05v-3q3.8.375 6.4 3.2T22 12q0 1.1-.2 2.125t-.65 1.925Z"/>`),
		g.Group(children),
	)
}