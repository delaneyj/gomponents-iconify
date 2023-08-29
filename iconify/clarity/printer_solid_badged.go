package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33 12.88a7.3 7.3 0 0 1-4 .55V15h-2v-2h.32a7.52 7.52 0 0 1-4.18-4H11V7h11.57a7.52 7.52 0 0 1-.07-1a7.54 7.54 0 0 1 .07-1H9v4H7a4 4 0 0 0-4 4v11h6v8h18v-8h6V12.88ZM25 24v6H11V19h14Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}