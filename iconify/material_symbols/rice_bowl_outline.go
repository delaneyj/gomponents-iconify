package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RiceBowlOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22v-1.75Q5.375 19.2 3.687 17T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.8-1.688 5T16 20.25V22H8Zm2-10h4V4.25q-.5-.125-1-.188T12 4q-.5 0-1 .063t-1 .187V12Zm-6 0h4V5.075Q6.125 6.15 5.062 8T4 12Zm12 0h4q0-2.15-1.063-4T16 5.075V12Zm-6 8h4v-1.1q1.8-.725 3.325-1.913T19.6 14H4.4q.75 1.8 2.275 2.988T10 18.9V20Zm0 0h4h-4Z"/>`),
		g.Group(children),
	)
}