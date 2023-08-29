package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignNoRightTurn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0Zm-3.416 5.29L6.596 7.304A1.498 1.498 0 0 0 6 8.5V11H5V8.5c0-.765.344-1.45.885-1.908L2.709 3.416a7 7 0 0 0 9.874 9.874Zm.707-.706A7 7 0 0 0 3.417 2.71l3.388 3.388C7.025 6.034 7.259 6 7.5 6H9V4.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L9.41 8.658a.265.265 0 0 1-.026.02l3.907 3.906ZM7.707 7L9 8.293V7H7.707Z"/>`),
		g.Group(children),
	)
}