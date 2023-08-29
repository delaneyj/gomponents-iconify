package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsPhotoCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.25 22l-.4-3.2q-.325-.125-.613-.3t-.562-.375L4.7 19.375l-2.75-4.75l2.575-1.95Q4.5 12.5 4.5 12.337v-.674q0-.163.025-.338L1.95 9.375l2.75-4.75l2.975 1.25q.275-.2.575-.375t.6-.3l.4-3.2h5.5l.4 3.2q.325.125.613.3t.562.375l2.975-1.25l2.75 4.75L19.925 11H15.4q-.35-1.075-1.25-1.788t-2.1-.712q-1.45 0-2.475 1.025T8.55 12q0 1.2.675 2.1T11 15.35V22H9.25Zm5.25 0q-.625 0-1.063-.438T13 20.5v-5q0-.625.438-1.063T14.5 14H16l1-1h2l1 1h1.5q.625 0 1.063.438T23 15.5v5q0 .625-.438 1.063T21.5 22h-7Zm3.5-2q.825 0 1.413-.588T20 18q0-.825-.588-1.413T18 16q-.825 0-1.413.588T16 18q0 .825.588 1.413T18 20Z"/>`),
		g.Group(children),
	)
}