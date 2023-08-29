package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17q-.425 0-.7-.3t-.225-.7q.35-2.55 2.313-4.275T12 10q2.65 0 4.612 1.725T18.925 16q.05.4-.225.7t-.7.3q-.4 0-.7-.275t-.4-.7q-.35-1.725-1.712-2.875T12 12q-1.825 0-3.188 1.15T7.1 16.025q-.1.425-.4.7T6 17Zm-4 0q-.425 0-.713-.3T1.05 16q.2-2.1 1.125-3.925T4.563 8.9q1.462-1.35 3.375-2.125T12 6q2.15 0 4.063.775T19.438 8.9q1.462 1.35 2.4 3.175T22.95 16q.05.4-.237.7T22 17q-.4 0-.7-.288T20.95 16q-.375-3.4-2.912-5.7T12 8q-3.5 0-6.038 2.3T3.05 16q-.05.425-.337.713T2 17Z"/>`),
		g.Group(children),
	)
}