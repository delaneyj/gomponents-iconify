package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screenshottile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 5.5h-9a2 2 0 0 0-2 2v9m37 0v-9a2 2 0 0 0-2-2h-9m-26 26v9a2 2 0 0 0 2 2h9m15 0h9a2 2 0 0 0 2-2v-9"/><circle cx="24" cy="24" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}