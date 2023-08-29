package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouponFourLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.005 21h-7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h7a2 2 0 1 0 4 0h7a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-7a2 2 0 1 0-4 0ZM8.54 19a3.998 3.998 0 0 1 3.465-2c1.48 0 2.773.804 3.465 2h4.535V5H15.47a3.999 3.999 0 0 1-3.465 2A3.998 3.998 0 0 1 8.54 5H4.005v14H8.54ZM6.005 8h2v8h-2V8Zm10 0h2v8h-2V8Z"/>`),
		g.Group(children),
	)
}