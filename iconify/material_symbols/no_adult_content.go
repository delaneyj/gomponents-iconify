package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoAdultContent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-2q1.4 0 2.638-.438t2.262-1.237L13.575 15H16.4l1.9 1.9q.8-1.025 1.25-2.263T20 12q0-3.325-2.337-5.663T12 4q-1.4 0-2.638.45T7.1 5.7L10.425 9H7.6L5.7 7.1q-.8 1.025-1.25 2.263T4 12q0 3.325 2.337 5.663T12 20Zm-7-6l1.5-2L5 10h1.5l.75 1L8 10h1.5L8 12l1.5 2H8l-.75-1l-.75 1H5Zm4.75 0l1.5-2l-1.5-2h1.5l.75 1l.75-1h1.5l-1.5 2l1.5 2h-1.5L12 13l-.75 1h-1.5Zm4.75 0l1.5-2l-1.5-2H16l.75 1l.75-1H19l-1.5 2l1.5 2h-1.5l-.75-1l-.75 1h-1.5Z"/>`),
		g.Group(children),
	)
}