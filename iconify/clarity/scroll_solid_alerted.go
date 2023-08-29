package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrollSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28.08 15.4v14.14a2.5 2.5 0 0 1-2.5 2.5H10.24a4.47 4.47 0 0 0 .76-2.5v-23a2.5 2.5 0 0 1 5 0v4.54h2.61A3.66 3.66 0 0 1 19 9.89l4.51-7.8H13.5A4.5 4.5 0 0 0 9 6.58V24H2v5.5A4.5 4.5 0 0 0 6.5 34h19.08a4.5 4.5 0 0 0 4.5-4.5V15.4Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}