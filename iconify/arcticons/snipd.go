package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snipd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.9 33.1c-6.6-6.6-6.6-17.2 0-23.7s17.2-6.6 23.7 0L14.9 33.1ZM33.1 15c6.6 6.6 6.6 17.2 0 23.7s-17.2 6.6-23.7 0L33.1 15Z"/>`),
		g.Group(children),
	)
}