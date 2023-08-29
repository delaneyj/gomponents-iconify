package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FogSunny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1v3h-2V1h2Zm7.485 3.928L18.364 7.05L16.95 5.636l2.121-2.122l1.414 1.414ZM4.93 3.514l2.12 2.122L5.636 7.05L3.515 4.929l1.414-1.415ZM12 8a4 4 0 0 0-3.668 5.6l.4.916l-1.832.8l-.4-.916a6 6 0 1 1 11 0l-.4.917l-1.833-.801l.4-.916A4 4 0 0 0 12 8ZM1 11h3v2H1v-2Zm19 0h3v2h-3v-2ZM2 17h6v2H2v-2Zm8 0h12v2H10v-2Zm7 4h5v2h-5v-2ZM2 21h13v2H2v-2Z"/>`),
		g.Group(children),
	)
}