package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viaplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5a21.37 21.37 0 0 0-13.75 5L32 19.12v10L10.3 40.57A21.5 21.5 0 1 0 24 2.5Zm-19.28 12a21.48 21.48 0 0 0 0 19l18-9.51Z"/>`),
		g.Group(children),
	)
}