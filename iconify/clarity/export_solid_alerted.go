package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExportSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29.32 16.35a1 1 0 0 0-1.41 1.41L31.16 21H26v2h5.19l-3.28 3.28a1 1 0 1 0 1.41 1.41L35 22Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M17 22a1 1 0 0 1 1-1h8v-5.6h-3.77A3.68 3.68 0 0 1 19 9.89L22.45 4H10.87L4 10.86V30a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2v-7h-8a1 1 0 0 1-1-1Zm-5-10H6v-.32L11.69 6H12Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-3--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}