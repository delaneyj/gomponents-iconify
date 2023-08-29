package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDecreaseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19L5.975 5.75Q6.1 5.4 6.5 5.2T7.95 5q.375 0 .663.2t.412.55L13.45 17.5q.2.55-.138 1.025T12.4 19q-.35 0-.65-.2t-.425-.55l-.375-1.05q-.275-.8-1.2-1.288t-3.2-.487q-.85 0-1.537.488T4.05 17.2L3.4 19H1Zm4.4-5.6h4.2L7.55 7.6h-.1L5.4 13.4ZM16 13q-.425 0-.713-.288T15 12q0-.425.288-.713T16 11h6q.425 0 .713.288T23 12q0 .425-.288.713T22 13h-6Z"/>`),
		g.Group(children),
	)
}