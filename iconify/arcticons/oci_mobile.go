package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OciMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.499 8.5c7.18 0 13 6.94 13 15.5h0c0 8.56-5.82 15.5-13 15.5H17.5c-4.644 0-8.936-2.954-11.258-7.75a18.114 18.114 0 0 1 0-15.5C8.564 11.454 12.856 8.5 17.5 8.5Z"/>`),
		g.Group(children),
	)
}