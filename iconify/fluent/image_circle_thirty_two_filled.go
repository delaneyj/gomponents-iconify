package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCircleThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.228 26.529A13.946 13.946 0 0 1 16 30c-3.535 0-6.764-1.31-9.228-3.471l8.874-8.882a.5.5 0 0 1 .708 0l8.874 8.882Zm1.406-1.423l-8.865-8.873a2.5 2.5 0 0 0-3.538 0l-8.865 8.873A13.945 13.945 0 0 1 2 16C2 8.268 8.268 2 16 2s14 6.268 14 14c0 3.477-1.268 6.658-3.366 9.106ZM24 11.25a3.25 3.25 0 1 0-6.5 0a3.25 3.25 0 0 0 6.5 0Z"/>`),
		g.Group(children),
	)
}