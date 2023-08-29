package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.91 2.602a3.25 3.25 0 0 0-3.82 0L3.847 10.043a3.25 3.25 0 0 0-1.18 3.634l3.911 12.04a3.25 3.25 0 0 0 3.091 2.245h12.66a3.25 3.25 0 0 0 3.09-2.245l3.912-12.04a3.25 3.25 0 0 0-1.18-3.634L17.91 2.603Z"/>`),
		g.Group(children),
	)
}