package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResourcePoolOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33.68 15.4h-1.95a14 14 0 0 1 .22 1.6H17.49L8.3 28.07A14 14 0 0 1 22.09 4.62l1-1.76A16 16 0 1 0 34 18a16 16 0 0 0-.23-2.61ZM18 32a13.91 13.91 0 0 1-8.16-2.65L18.43 19h13.52A14 14 0 0 1 18 32Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}