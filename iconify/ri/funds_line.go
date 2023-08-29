package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FundsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.41 14.526l3.402-3.402l2.829 2.829l3.157-3.157l-1.793-1.793h5v5l-1.793-1.793l-4.571 4.571l-2.829-2.828l-2.474 2.474a8 8 0 1 0-.927-1.9Zm-1.537 1.558l-.01-.01l.004-.004a9.965 9.965 0 0 1-.862-4.067c0-5.523 4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10c-4.07 0-7.57-2.43-9.132-5.919Z"/>`),
		g.Group(children),
	)
}