package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.587 12.074a.765.765 0 0 0-.188-.202a.75.75 0 1 0-.881 1.214c.302.22.707.167.965-.09L6.4 22.762L4.195 30.76l6.917-4.577L26.865 4.468l-4.716-3.42l-1.52 2.096a1.588 1.588 0 0 0-.597-.907a1.63 1.63 0 0 0-2.28.363c-3.032 4.182-1.35 5.296-4.166 9.474zm-3.47 13.074L6.56 27.503l1.133-4.117l2.425 1.762zM14.31 11.86c2.182-3.224 1.974-4.098 3.842-6.96c.31.21.664.286 1.012.268L14.31 11.86z"/>`),
		g.Group(children),
	)
}