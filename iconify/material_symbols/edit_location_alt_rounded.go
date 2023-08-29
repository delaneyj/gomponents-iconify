package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditLocationAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.325q-.35 0-.7-.125t-.625-.375Q9.05 19.325 7.8 17.9t-2.087-2.762q-.838-1.338-1.275-2.575T4 10.2q0-3.75 2.413-5.975T12 2q.675 0 1.338.113t1.287.312L9.575 7.5q-.275.275-.425.638T9 8.9V11q0 .825.588 1.412T11 13h2.1q.4 0 .763-.15t.637-.425l5.05-5.05q.225.65.337 1.35T20 10.2q0 2.35-1.7 5.037t-4.975 5.588q-.275.25-.625.375t-.7.125ZM18.35 3.85l.95.95l-5.9 5.9q-.15.15-.338.225t-.387.075h-.925q-.325 0-.537-.212T11 10.25v-.925q0-.2.075-.388T11.3 8.6l5.9-5.9l1.15 1.15ZM20 4.1L17.9 2l.7-.7q.275-.275.7-.275t.7.275l.7.7q.275.275.275.7t-.275.7l-.7.7Z"/>`),
		g.Group(children),
	)
}