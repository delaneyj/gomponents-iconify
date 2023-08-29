package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WakingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.028 12.847l-.005.003a12.587 12.587 0 1 0-21.726 12.041l-.004.002l.026.027a25.203 25.203 0 0 0 1.609 1.805C18.032 30.983 30.242 43.5 30.242 43.5V30.463h5.303V23h2.873Z"/>`),
		g.Group(children),
	)
}