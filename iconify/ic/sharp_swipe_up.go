package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSwipeUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.06 5.56L1 4.5L4.5 1L8 4.5L6.94 5.56L5.32 3.94a10.457 10.457 0 0 0 1.88 8.99L6.13 14A11.974 11.974 0 0 1 3.5 6.5c0-.92.1-1.82.3-2.68L2.06 5.56zm19.65 5.62l2.09 7.39l-8.23 3.65l-6.84-2.85l.61-1.62l3.8-.75l-4.35-9.83c-.34-.76 0-1.64.76-1.98c.76-.34 1.64 0 1.98.76l2.43 5.49l1.26-.56l6.49.3z"/>`),
		g.Group(children),
	)
}