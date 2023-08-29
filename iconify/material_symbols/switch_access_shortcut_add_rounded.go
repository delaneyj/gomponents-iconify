package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchAccessShortcutAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 14h-1q-.425 0-.713-.288T16 13q0-.425.288-.713T17 12h1v-1q0-.425.288-.713T19 10q.425 0 .713.288T20 11v1h1q.425 0 .713.288T22 13q0 .425-.288.713T21 14h-1v1q0 .425-.288.713T19 16q-.425 0-.713-.288T18 15v-1Zm-5-8.7q-1.425 1.275-2.212 2.988T10 11.9q0 2.275 1.088 4.263t3.037 3.187q.35.225.488.613t-.063.762q-.2.375-.575.513t-.725-.063Q10.8 19.7 9.4 17.237T8 11.9q0-2.275.9-4.313T11.45 4H9q-.425 0-.713-.288T8 3q0-.425.288-.713T9 2h5q.425 0 .713.288T15 3v5q0 .425-.288.713T14 9q-.425 0-.713-.288T13 8V5.3Z"/>`),
		g.Group(children),
	)
}