package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28 18H8a1 1 0 0 0 0 2h1v12h18V20h1a1 1 0 0 0 0-2Zm-3 12H11V20h14Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M27 13.04h2v2h-2z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M33 12.88a7.45 7.45 0 0 1-2 .55v8.66h-1.92V24H33V12.88Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M5 13a2 2 0 0 1 2-2h17.42a7.5 7.5 0 0 1-1.27-2H11V7h11.57a7.52 7.52 0 0 1-.07-1a7.54 7.54 0 0 1 .07-1H9v4H7a4 4 0 0 0-4 4v11h3.92v-1.91H5Z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-5--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}