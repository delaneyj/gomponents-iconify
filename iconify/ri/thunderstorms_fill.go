package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThunderstormsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.988 18l1.216-1.58a1.5 1.5 0 0 0-1.189-2.415H15v-3.976a1.5 1.5 0 0 0-2.69-.914l-6.365 8.281A8.002 8.002 0 0 1 9 2a8.003 8.003 0 0 1 7.458 5.099A5.5 5.5 0 1 1 17.5 18h-.512ZM13 16.005h3l-5 6.5v-4.5H8l5-6.505v4.505Z"/>`),
		g.Group(children),
	)
}