package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetForwardCircleF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.514 11.836V14a1 1 0 0 0 2 0V6a1 1 0 1 0-2 0v2.164L8.585 5.356A1.887 1.887 0 0 0 7.487 5c-1.09 0-1.973.941-1.973 2.102v5.796c0 .417.117.824.335 1.17c.606.965 1.831 1.222 2.736.576l3.93-2.808zM10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zM7.355 7.08l4.055 2.898l-4.055 2.898V7.08z"/>`),
		g.Group(children),
	)
}