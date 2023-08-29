package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M22.23 15.4A3.68 3.68 0 0 1 19 9.89l.54-.89H11V7h9.71l1.13-2H9v4H7a4 4 0 0 0-4 4v11h6v8h18v-8h6v-8.6ZM25 24v6H11V19h14Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}