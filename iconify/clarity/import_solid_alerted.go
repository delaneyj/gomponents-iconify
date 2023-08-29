package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M3 21a1 1 0 1 0 0 2h5v-2Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M22.23 15.4A3.68 3.68 0 0 1 19 9.89L22.45 4h-7.58L8 10.86V21h7.2l-3.25-3.25a1 1 0 0 1 1.41-1.41L19 22l-5.68 5.68a1 1 0 0 1-1.41-1.41L15.23 23H8v7a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V15.4ZM16 12h-6v-.32L15.69 6H16Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-3--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}