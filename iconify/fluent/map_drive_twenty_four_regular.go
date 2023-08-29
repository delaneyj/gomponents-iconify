package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapDriveTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.998 8.002a1.002 1.002 0 1 0 0-2.004a1.002 1.002 0 0 0 0 2.004ZM4.25 3A2.25 2.25 0 0 0 2 5.25v5c0 .414.336.75.75.75h8.5v2.001h-1a.75.75 0 0 0-.75.75v2.254H7.752a.75.75 0 0 0-.75.75v.745H2.75a.75.75 0 0 0 0 1.5h4.251v1.253c0 .415.336.75.75.75h8.498a.75.75 0 0 0 .75-.75V19h4.251a.75.75 0 0 0 0-1.5H17v-.745a.75.75 0 0 0-.75-.75H14.5V13.75a.75.75 0 0 0-.75-.75h-1v-2h8.5a.75.75 0 0 0 .75-.75v-5A2.25 2.25 0 0 0 19.75 3H4.25ZM3.5 5.25a.75.75 0 0 1 .75-.75h15.5a.75.75 0 0 1 .75.75V9.5h-17V5.25Zm7.5 9.251h2v2.254c0 .414.335.75.75.75h1.75v1.998h-7v-1.998h1.75a.75.75 0 0 0 .75-.75V14.5Z"/>`),
		g.Group(children),
	)
}