package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpRightTwoShortArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M20.991 3.009L3.001 21m16.93-2.081c-.786-.786-1.183-2.737-1.382-4.51c-.258-2.282-.159-4.6.381-6.834c.404-1.673 1.056-3.543 2.07-4.557M5.081 4.07c.786.785 2.737 1.182 4.51 1.381c2.282.258 4.6.159 6.834-.381c1.673-.404 3.543-1.056 4.557-2.07"/>`),
		g.Group(children),
	)
}