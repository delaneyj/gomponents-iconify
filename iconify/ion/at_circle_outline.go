package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32" d="M448.08 256.08c0-106-86-192-192-192s-192 86-192 192s86 192 192 192s192-85.97 192-192Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="28" d="M300.81 358.29c-20.83 7.42-34.05 9.59-54.19 9.59c-61.17 0-106.39-50.07-101-111.84S205 144.21 266.14 144.21c68.92 0 106.79 45.55 101.47 106.55c-4 45.54-32.8 58.66-47.89 56c-14.2-2.55-25.92-15.52-23.75-40.35l5.62-44.66c-7.58-9.17-28.11-18-49.93-14.54C231.77 210.3 209 228 206.56 256s14.49 50.84 39.93 50.84s47.86-18.39 50.69-50.84"/>`),
		g.Group(children),
	)
}