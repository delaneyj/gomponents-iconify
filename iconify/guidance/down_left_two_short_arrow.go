package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftTwoShortArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.01 20.991L21 3.001M4.07 5.08c.785.786 1.182 2.737 1.381 4.51c.258 2.282.159 4.6-.381 6.834c-.404 1.673-1.056 3.543-2.07 4.557m15.92-1.051c-.785-.785-2.737-1.181-4.51-1.381c-2.282-.258-4.6-.159-6.834.381c-1.673.404-3.543 1.056-4.557 2.07"/>`),
		g.Group(children),
	)
}