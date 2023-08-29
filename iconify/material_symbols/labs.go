package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Labs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17V8q-.825 0-1.413-.588T5 6V4q0-.825.588-1.413T7 2h10q.825 0 1.413.588T19 4v2q0 .825-.588 1.413T17 8v9q0 2.075-1.463 3.538T12 22Zm0-2q.975 0 1.75-.563T14.825 18H12v-2h3v-1h-3v-2h3v-1h-3v-2h3V8H9v9q0 1.25.875 2.125T12 20Z"/>`),
		g.Group(children),
	)
}