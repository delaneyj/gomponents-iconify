package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 22.6l-3.85-3.85L15 22H9l-1.35-4.525q-1.2-.95-1.925-2.375T5 12q0-.9.225-1.737t.65-1.588L1.4 4.2l1.425-1.425l18.4 18.4L19.8 22.6ZM12 17q.5 0 .95-.088t.875-.262L7.35 10.175q-.175.425-.262.875T7 12q0 2.075 1.463 3.538T12 17Zm6.175-1.7l-1.5-1.5q.175-.425.25-.863T17 12q0-2.075-1.463-3.538T12 7q-.5 0-.938.075t-.862.25l-2.15-2.15L9 2h6l1.35 4.525q1.2.95 1.925 2.375T19 12q0 .9-.2 1.725t-.625 1.575Z"/>`),
		g.Group(children),
	)
}