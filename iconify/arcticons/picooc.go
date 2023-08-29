package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picooc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.5 42.5l-16-16m16-21l-16 16m-21 21l15.672-15.672a4 4 0 0 0 0-5.656L5.5 5.5"/>`),
		g.Group(children),
	)
}