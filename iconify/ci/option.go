package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Option(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7h5.094c.33 0 .495 0 .643.047c.132.042.253.111.357.202c.117.103.202.245.372.528l5.068 8.446c.17.284.255.425.372.528c.103.09.224.16.356.202c.148.047.314.047.644.047H21M15 7h6"/>`),
		g.Group(children),
	)
}