package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 14q-.825 0-1.413-.588T0 12V6q0-.3.125-.575T.45 4.95L5.4 0l.75.75q.15.15.25.388t.1.462v.2L5.8 5H11q.425 0 .713.288T12 6v1.25q0 .15-.025.288T11.9 7.8l-2.25 5.3q-.175.425-.563.663T8.25 14H2Zm5.95-2L10 7.15V7H3.35l.6-2.7L2 6.2V12h5.95ZM18.6 24l-.75-.75q-.15-.15-.25-.388t-.1-.462v-.2l.7-3.2H13q-.425 0-.713-.288T12 18v-1.25q0-.15.025-.288t.075-.262l2.25-5.3q.2-.425.575-.663T15.75 10H22q.825 0 1.413.587T24 12v6q0 .3-.113.563t-.337.487L18.6 24Zm-2.55-12L14 16.85V17h6.65l-.6 2.7L22 17.8V12h-5.95ZM2 12V6.2l1.95-1.9l-.6 2.7H10v.15L7.95 12H2Zm20 0v5.8l-1.95 1.9l.6-2.7H14v-.15L16.05 12H22Z"/>`),
		g.Group(children),
	)
}