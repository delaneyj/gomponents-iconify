package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.783 2.06C7.25 1.527 9.365 1 12 1c2.634 0 4.75.527 6.217 1.06a14.13 14.13 0 0 1 1.699.74a9.315 9.315 0 0 1 .62.356l.011.007l.005.003l.001.001l.002.001l.445.297V20h1v2H2v-2h1V3.465l.445-.297l.002-.001l.001-.001l.005-.003l.011-.007a3.4 3.4 0 0 1 .163-.101c.107-.064.26-.151.457-.254c.395-.206.966-.474 1.7-.74ZM5 4.58V20h2V7.5h2v2h2v-2h2V20h6V4.579l-.01-.005c-.324-.17-.815-.4-1.457-.634C16.25 3.473 14.365 3 12 3s-4.25.473-5.533.94A12.142 12.142 0 0 0 5 4.579Zm4 6.92V13h2v-1.5H9Zm2 3.5H9v1.5h2V15Zm0 3.5H9V20h2v-1.5Z"/>`),
		g.Group(children),
	)
}