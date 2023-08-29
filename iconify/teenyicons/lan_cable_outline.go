package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanCableOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 11.5V15m2-3.5V15M6 9.5h3M6.5.5v2h2v-2m-4 0h6v4h1v3l-2 4h-4l-2-4v-3h1v-4Zm2 4v2a1 1 0 0 0 2 0v-2h-2Z"/>`),
		g.Group(children),
	)
}