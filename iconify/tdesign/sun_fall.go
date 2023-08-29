package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunFall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3v3h-2V3h2Zm7.485 3.928L18.364 9.05L16.95 7.636l2.121-2.122l1.414 1.414ZM4.93 5.514l2.12 2.122L5.636 9.05L3.515 6.929l1.414-1.415ZM12 10a4 4 0 0 0-4 4v1H6v-1a6 6 0 0 1 12 0v1h-2v-1a4 4 0 0 0-4-4ZM1 13h3v2H1v-2Zm19 0h3v2h-3v-2ZM1 17h8.303L12 18.798L14.697 17H23v2h-7.697L12 21.202L8.697 19H1v-2Z"/>`),
		g.Group(children),
	)
}