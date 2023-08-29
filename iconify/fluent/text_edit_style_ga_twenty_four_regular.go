package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextEditStyleGaTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M10.75 2.5a.75.75 0 0 1 .75.75V6.5H13A.75.75 0 0 1 13 8h-1.5v3.03L10 12.529V3.25a.75.75 0 0 1 .75-.75zm-.425 11.118l1.175-1.174l4.009-4.001a3.22 3.22 0 0 1 4.553.002a3.218 3.218 0 0 1-.001 4.551l-7.123 7.112a2.25 2.25 0 0 1-.943.562L7.702 21.96a1 1 0 0 1-1.24-1.264l1.362-4.229c.11-.34.299-.65.552-.902l1.949-1.946zm-.89 3.007a.75.75 0 0 0-.183.301l-1.07 3.323l3.382-1.015a.75.75 0 0 0 .314-.188L19 11.936a1.718 1.718 0 1 0-2.431-2.432l-7.133 7.121zM2 4.25a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 .747.682c.194 2.135-1.141 5.97-5.531 7.286a.75.75 0 1 1-.431-1.436C5.667 9.092 6.833 6.634 6.995 5H2.75A.75.75 0 0 1 2 4.25z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}