package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonApron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12q-1.65 0-2.825-1.175T8 8q0-1.65 1.175-2.825T12 4q1.65 0 2.825 1.175T16 8q0 1.65-1.175 2.825T12 12Zm4 8v-6.4q.625.2 1.225.425t1.175.525q.75.375 1.175 1.088T20 17.2V20h-4Zm-6-3.5v-3.35q.5-.075 1-.113T12 13q.5 0 1 .038t1 .112v3.35h-4ZM4 20v-2.8q0-.85.425-1.563T5.6 14.55q.575-.3 1.175-.525T8 13.6V20H4Z"/>`),
		g.Group(children),
	)
}