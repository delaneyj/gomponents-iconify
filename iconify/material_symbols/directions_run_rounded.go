package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsRunRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 5.5q-.825 0-1.413-.588T11.5 3.5q0-.825.588-1.413T13.5 1.5q.825 0 1.413.588T15.5 3.5q0 .825-.588 1.413T13.5 5.5ZM14 23q-.425 0-.713-.288T13 22v-5l-2.1-2l-.575 2.5q-.2.8-.875 1.238T7.975 19L4 18.2q-.425-.075-.663-.425T3.2 17q.075-.425.425-.663T4.4 16.2l3.8.8l1.6-8.1l-1.8.7V12q0 .425-.288.713T7 13q-.425 0-.713-.288T6 12V9.625q0-.6.325-1.113t.875-.737L9.95 6.6q.875-.375 1.288-.488T12 6q.525 0 .975.275T13.7 7l1 1.6q.55.875 1.438 1.5t2.012.825q.375.05.613.325t.237.625q0 .475-.338.8t-.787.25q-1.325-.2-2.45-.838T13.5 10.5l-.6 3l1.475 1.4q.3.3.463.663t.162.787V22q0 .425-.288.713T14 23Z"/>`),
		g.Group(children),
	)
}