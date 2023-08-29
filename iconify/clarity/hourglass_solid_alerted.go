package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28.67 32h-22a1 1 0 0 0 0 2h22a1 1 0 1 0 0-2Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M6.67 4h15.78l1.15-2H6.67a1 1 0 1 0 0 2Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted"/><path fill="currentColor" d="M12.51 20.27a6.07 6.07 0 0 0-2.45 4.55v5.12H25v-5.12a6.07 6.07 0 0 0-2.45-4.55a11.48 11.48 0 0 0-2.91-1.72v-1.16a11.48 11.48 0 0 0 2.91-1.72l.3-.27h-.62A3.68 3.68 0 0 1 19 9.89L21.29 6H10.06v5.12a6.07 6.07 0 0 0 2.45 4.55a11.48 11.48 0 0 0 2.91 1.72v1.16a11.48 11.48 0 0 0-2.91 1.72Z" class="clr-i-solid--alerted clr-i-solid-path-3--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-4--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}