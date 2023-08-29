package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotAccessible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.775 22.6l-5.6-5.6H12q-.825 0-1.413-.588T10 15v-2.175l-8.6-8.6L2.8 2.8l18.4 18.4l-1.425 1.4ZM10 22q-2.075 0-3.537-1.463T5 17q0-1.825 1.137-3.188T9 12.1v2.075q-.875.325-1.438 1.088T7 17q0 1.25.875 2.125T10 20q.975 0 1.75-.563T12.825 18H14.9q-.35 1.725-1.713 2.863T10 22Zm2-16q-.825 0-1.412-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Zm7 7q-1.325 0-2.675-.575T14 11.05l-3.425-3.425q.25-.275.575-.438t.85-.162q.375 0 .825.175t.825.55l1.275 1.425q.725.8 1.812 1.325T19 11v2Z"/>`),
		g.Group(children),
	)
}