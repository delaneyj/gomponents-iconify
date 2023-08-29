package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextEditStyleGaTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M10.75 2.5a.75.75 0 0 1 .75.75V6.5H13A.75.75 0 0 1 13 8h-1.5v3.03L10 12.529V3.25a.75.75 0 0 1 .75-.75zm-.425 11.118l1.175-1.174v.001l4.01-4.002a3.22 3.22 0 0 1 4.553.002a3.218 3.218 0 0 1-.002 4.551l-7.114 7.102a2.25 2.25 0 0 1-.978.573l-4.613 1.303a.75.75 0 0 1-.92-.94l1.387-4.543c.107-.354.3-.675.562-.936l1.94-1.937zM2 4.25a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 .747.682c.194 2.135-1.141 5.97-5.531 7.286a.75.75 0 1 1-.431-1.436C5.667 9.092 6.833 6.634 6.995 5H2.75A.75.75 0 0 1 2 4.25z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}