package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhatshotRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.5 0-4.588-1.1T3.95 17.95l4.1-4.1l2.3 1.925q.3.25.688.225t.662-.3l4.3-4.3V13q0 .425.288.712T17 14q.425 0 .713-.288T18 13V9q0-.425-.288-.713T17 8h-4q-.425 0-.713.288T12 9q0 .425.288.713T13 10h1.6l-3.65 3.65l-2.3-1.925q-.3-.25-.688-.225t-.662.3l-4.4 4.4q-.425-.95-.662-2.013T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}