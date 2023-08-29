package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28 18H8a1 1 0 0 0 0 2h1v12h18V20h1a1 1 0 0 0 0-2Zm-3 12H11V20h14Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M31 15.4v6.69h-1.92V24H33v-8.6h-2z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="M5 13a2 2 0 0 1 2-2h11.64A3.65 3.65 0 0 1 19 9.89l.54-.89H11V7h9.71l1.13-2H9v4H7a4 4 0 0 0-4 4v11h3.92v-1.91H5Z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-4--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}