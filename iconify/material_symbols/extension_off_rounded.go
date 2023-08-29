package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExtensionOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.075 21.9l-.925-.9H14.2q0-1.25-.787-2.125T11.5 18q-1.125 0-1.913.875T8.8 21H5q-.825 0-1.413-.588T3 19v-3.8q1.25 0 2.125-.787T6 12.5q0-1.125-.875-1.912T3 9.8V5.85l-.925-.925q-.275-.275-.275-.7t.3-.725q.275-.275.7-.275t.725.275l16.975 17q.275.275.275.7t-.275.7q-.3.3-.725.3t-.7-.3ZM20 17.15L6.85 4H9q0-1.05.725-1.775T11.5 1.5q1.05 0 1.775.725T14 4h4q.825 0 1.413.588T20 6v4q1.05 0 1.775.725T22.5 12.5q0 1.05-.725 1.775T20 15v2.15Z"/>`),
		g.Group(children),
	)
}