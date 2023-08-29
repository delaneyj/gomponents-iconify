package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonAddDisabledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.1 21.875L16.8 19.6q-.15.2-.35.3T16 20H2q-.425 0-.713-.288T1 19v-1.8q0-.85.438-1.563T2.6 14.55q1.55-.775 3.15-1.163T9 13q.325 0 .638.013t.612.037L9.2 12H9q-1.65 0-2.825-1.175T5 8v-.2L2.1 4.9q-.3-.3-.288-.7t.313-.7q.3-.3.713-.3t.712.3l16.975 16.975q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3ZM19 14q-.425 0-.713-.288T18 13v-2h-2q-.425 0-.713-.288T15 10q0-.425.288-.713T16 9h2V7q0-.425.288-.713T19 6q.425 0 .713.288T20 7v2h2q.425 0 .713.288T23 10q0 .425-.288.713T22 11h-2v2q0 .425-.288.713T19 14Zm-6.4-4.3L7.3 4.4q.375-.2.813-.3T9 4q1.65 0 2.825 1.175T13 8q0 .45-.1.888t-.3.812Z"/>`),
		g.Group(children),
	)
}